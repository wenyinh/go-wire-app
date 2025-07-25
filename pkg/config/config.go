package config

type AppConfiguration struct {
	DBConfig    DatabaseConfig `mapstructure:"database"`
	AppConfig   AppConfig      `mapstructure:"app"`
	RedisConfig RedisConfig    `mapstructure:"redis"`
}
