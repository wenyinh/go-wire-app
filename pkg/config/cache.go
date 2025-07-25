package config

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`     // e.g. "localhost:6379"
	Password string `mapstructure:"password"` // "" if no auth
	DB       int    `mapstructure:"db"`       // redis db index
}
