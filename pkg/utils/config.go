package utils

import (
	"flag"
	"fmt"
	"github.com/wenyinh/go-wire-app/pkg/config"
	"github.com/wenyinh/go-wire-app/pkg/constants"
	"os"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

func LoadConfiguration() (*config.AppConfiguration, error) {
	var configType string
	var configFile string
	f := flag.NewFlagSet(constants.AppName, flag.ExitOnError)
	f.StringVar(&configType, "config-type", "yaml", "the type of config file")
	f.StringVar(&configFile, "config-file", "./application.yaml", "the location of config file")
	err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Failed to parse command flags, error:", err)
		os.Exit(1)
	}
	return LoadConfigurationFromFile(configFile, configType)
}

func LoadConfigurationFromFile(configFile, configType string) (*config.AppConfiguration, error) {
	var conf config.AppConfiguration
	// Set default value to configs
	defaults.SetDefaults(&conf)

	// Read configs from config file
	c := viper.New()
	c.SetConfigFile(configFile)
	c.SetConfigType(configType)
	if err := c.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config file, error:", err)
		return nil, err
	}
	if err := c.Unmarshal(&conf); err != nil {
		fmt.Println("Failed to parse config file, error:", err)
		return nil, err
	}

	return &conf, nil
}
