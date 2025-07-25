package config

import "time"

type DatabaseConfig struct {
	// options: mysql, sqlite
	Driver string `mapstructure:"driver"`
	// driver=mysql, example: username:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True
	// driver=sqlite, example: ./database.db
	Address           string        `mapstructure:"address"`
	MaxIdleConns      int           `mapstructure:"maxIdleConns" default:"10"`
	MaxOpenConns      int           `mapstructure:"maxOpenConns" default:"10"`
	ConnMaxLifetime   time.Duration `mapstructure:"connMaxLifetime" default:"1h"`
	DebugMode         bool          `mapstructure:"debugMode" default:"false"`
	SlowSQLThreshold  time.Duration `mapstructure:"slowSQLThreshold" default:"0.5s"`
	EnableAutoMigrate bool          `mapstructure:"enableAutoMigrate" default:"false"`
}
