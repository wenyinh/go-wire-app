package config

type AppConfig struct {
	AppName string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
}
