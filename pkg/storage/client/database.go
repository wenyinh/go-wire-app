package client

import (
	"errors"
	"fmt"
	"github.com/wenyinh/go-wire-app/pkg/config"
	"github.com/wenyinh/go-wire-app/pkg/storage/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	driverMysql  = "mysql"
	driverSqlite = "sqlite"
)

type DBClient interface {
	Database() *gorm.DB
	Close() error
}

type GormDBClient struct {
	db     *gorm.DB
	config *config.DatabaseConfig
}

func NewDatabaseClient(databaseConfig *config.DatabaseConfig) (*GormDBClient, func(), error) {
	db, err := makeGormDB(databaseConfig)
	if err != nil {
		return nil, nil, err
	}
	if databaseConfig.DebugMode {
		db = db.Debug()
	}
	connPool, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	connPool.SetMaxIdleConns(databaseConfig.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	connPool.SetMaxOpenConns(databaseConfig.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	connPool.SetConnMaxLifetime(databaseConfig.ConnMaxLifetime)

	client := &GormDBClient{
		db:     db,
		config: databaseConfig,
	}

	if client.config.EnableAutoMigrate {
		if pErr := autoMigrateDB(client); pErr != nil {
			return nil, nil, pErr
		}
	}

	return client, func() {
		_ = client.Close()
	}, nil
}

func makeGormDB(databaseConfig *config.DatabaseConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	gormConfig.Logger = NewCustomGormLogger(databaseConfig.SlowSQLThreshold)
	gormConfig.DisableForeignKeyConstraintWhenMigrating = true

	// Must be set to enable error comparison
	gormConfig.TranslateError = true

	switch databaseConfig.Driver {
	case driverMysql:
		return gorm.Open(mysql.Open(databaseConfig.Address), gormConfig)
	case driverSqlite:
		return gorm.Open(sqlite.Open(databaseConfig.Address), gormConfig)
	default:
		return nil, errors.New("unsupported database driver")
	}
}

func autoMigrateDB(client *GormDBClient) error {
	fmt.Println("[INFO] Database: enable auto migrate...")
	err := client.db.AutoMigrate(
		&model.UserDataModel{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *GormDBClient) Database() *gorm.DB {
	return c.db
}

func (c *GormDBClient) Close() error {
	connPool, err := c.db.DB()
	if err != nil {
		return err
	}
	return connPool.Close()
}
