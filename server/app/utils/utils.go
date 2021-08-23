package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Address      string `mapstructure:"Address"`
	Password     string `mapstructure:"Password"`
	DB           int    `mapstructure:"DB"`
	KeyName      string `mapstructure:"KeyName"`
	SearchLength int64  `mapstructure:"SearchLength"`
}

var Configuration Config

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
