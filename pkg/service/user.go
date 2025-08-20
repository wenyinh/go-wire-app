package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	goRedis "github.com/redis/go-redis/v9"
	dataModel "github.com/wenyinh/go-wire-app/pkg/storage/model"
	"github.com/wenyinh/go-wire-app/pkg/storage/redis"
	"github.com/wenyinh/go-wire-app/pkg/storage/repository"
	"github.com/wenyinh/go-wire-app/pkg/typed/entity"
	"github.com/wenyinh/go-wire-app/pkg/typed/param"
	"go.uber.org/zap"
	"strconv"
	"time"
)

/* ************************** Interface & Implementation ************************** */

type UserService interface {
	CreateUser(ctx context.Context, req param.CreateUserRequest) (*param.CreateUserResponse, error)
	GetUser(ctx context.Context, req param.GetUserRequest) (*param.GetUserResponse, error)
}

type UserServiceImpl struct {
	userRepository       repository.UserRepository
	cacheClientInterface redis.CacheClientInterface
}

// assertion
var _ UserService = &UserServiceImpl{}

func NewUserService(userRepository repository.UserRepository, cacheClientInterface redis.CacheClientInterface) (*UserServiceImpl, error) {
	return &UserServiceImpl{
		userRepository:       userRepository,
		cacheClientInterface: cacheClientInterface,
	}, nil
}

func (svc *UserServiceImpl) CreateUser(ctx context.Context, req param.CreateUserRequest) (*param.CreateUserResponse, error) {
	// 1. 构造 UserDataModel
	now := time.Now()
	userModel := &dataModel.UserDataModel{
		Username:   req.Username,
		Email:      req.Email,
		Gender:     req.Gender,
		Age:        req.Age,
		CreateTime: &now,
		UpdateTime: &now,
	}

	// 2. 插入数据库
	newModel, err := svc.userRepository.Create(ctx, userModel)
	if err != nil {
		return nil, err
	}
	// 3. 构造响应
	resp := &param.CreateUserResponse{
		UserId: newModel.ID,
	}
	return resp, nil
}

func (svc *UserServiceImpl) GetUser(ctx context.Context, req param.GetUserRequest) (*param.GetUserResponse, error) {
	logger := zap.L().With(zap.String("method", "UserService.GetUser"), zap.String("userID", strconv.FormatUint(req.UserId, 10)))
	cacheKey := fmt.Sprintf("user:%d", req.UserId)

	// 1. Ping Redis
	err := svc.cacheClientInterface.Ping(ctx)
	if err != nil {
		logger.Warn("failed to ping Redis", zap.Error(err))
		return nil, err
	}
	logger.Info("succeed to ping Redis")

	// 2. 尝试从 Redis 获取
	cached, err := svc.cacheClientInterface.RDB().Get(ctx, cacheKey).Result()
	if err == nil {
		var user entity.UserEntity
		unmarshalErr := json.Unmarshal([]byte(cached), &user)
		if unmarshalErr == nil {
			logger.Info("user fetched from Redis")
			return &param.GetUserResponse{User: user}, nil
		}
		logger.Warn("failed to unmarshal Redis user", zap.Error(unmarshalErr))
	} else if err != goRedis.Nil {
		logger.Warn("Redis Get error", zap.Error(err))
	}

	// 2. 查询数据库
	userModel, err := svc.userRepository.GetByID(ctx, req.UserId)
	if err != nil {
		logger.Error("failed to get user", zap.Error(err))
		return nil, err
	}
	if userModel == nil {
		logger.Warn("user not found")
		return nil, errors.New("user not found")
	}

	// 3. 转换并构造响应
	userEntity := entity.UserEntity{
		Username:   userModel.Username,
		Email:      userModel.Email,
		Gender:     userModel.Gender,
		Age:        userModel.Age,
		CreateTime: userModel.CreateTime,
		UpdateTime: userModel.UpdateTime,
	}

	resp := &param.GetUserResponse{
		User: userEntity,
	}

	// 4. 写入 Redis，过期时间 1 小时
	data, err := json.Marshal(userEntity)
	if err == nil {
		setErr := svc.cacheClientInterface.RDB().Set(ctx, cacheKey, data, time.Hour).Err()
		if setErr != nil {
			logger.Warn("failed to cache user in Redis", zap.Error(setErr))
		}
	} else {
		logger.Warn("failed to marshal user for Redis", zap.Error(err))
	}

	return resp, nil
}
