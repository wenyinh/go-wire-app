package repository

import (
	"context"
	"errors"
	"github.com/wenyinh/go-wire-app/pkg/constants"
	"github.com/wenyinh/go-wire-app/pkg/storage/client"
	"github.com/wenyinh/go-wire-app/pkg/storage/model"
	"github.com/wenyinh/go-wire-app/pkg/storage/query"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, model *model.UserDataModel) (*model.UserDataModel, error)
	DB() *gorm.DB
	Query(ctx context.Context) *query.Query
	GetByID(ctx context.Context, id uint64) (*model.UserDataModel, error)
}

type UserRepositoryImpl struct {
	dbClient client.DBClient
}

var _ UserRepository = &UserRepositoryImpl{}

func NewUserRepository(dbClient client.DBClient) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		dbClient: dbClient,
	}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, model *model.UserDataModel) (*model.UserDataModel, error) {
	q := repo.Query(ctx)
	userDO := q.UserDataModel.WithContext(ctx)
	err := userDO.Create(model)
	if err != nil {
		zap.L().Error("Failed to create user", zap.Error(err))
		return nil, err
	}
	return model, nil
}

func (repo *UserRepositoryImpl) GetByID(ctx context.Context, id uint64) (*model.UserDataModel, error) {
	q := repo.Query(ctx)
	userDO := q.UserDataModel.WithContext(ctx)

	user, err := userDO.Where(q.UserDataModel.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 查不到时返回 nil，不视为错误
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) DB() *gorm.DB {
	return repo.dbClient.Database()
}

func (repo *UserRepositoryImpl) Query(ctx context.Context) *query.Query {
	if q, ok := ctx.Value(constants.ContextKeyOfQueryTx).(*query.Query); ok {
		return q
	} else {
		return query.Use(repo.DB())
	}
}
