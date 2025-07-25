package app

import (
	"github.com/google/wire"
	"github.com/wenyinh/go-wire-app/api/v1/user"
	"github.com/wenyinh/go-wire-app/pkg/config"
	"github.com/wenyinh/go-wire-app/pkg/service"
	"github.com/wenyinh/go-wire-app/pkg/storage/client"
	"github.com/wenyinh/go-wire-app/pkg/storage/redis"
	"github.com/wenyinh/go-wire-app/pkg/storage/repository"
)

// Global
var wireSet = wire.NewSet(
	// Controller
	user.NewController, // Service
	// Service
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	service.NewUserService,
	// Repository
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	repository.NewUserRepository,
	// DBClient
	wire.Bind(new(client.DBClient), new(*client.GormDBClient)),
	client.NewDatabaseClient,
	// CacheClient
	wire.Bind(new(redis.CacheClientInterface), new(*redis.CacheClient)),
	redis.NewCacheClient,
	// Config
	wire.FieldsOf(
		new(*config.AppConfiguration),
		"DBConfig",
		"AppConfig",
		"RedisConfig",
	),
)
