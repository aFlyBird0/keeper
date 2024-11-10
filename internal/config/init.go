package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	S3 S3
	DB DB
}

var cfg Config

func Load() Config {
	return cfg
}

func init() {
	initViper()
}

func initViper() {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read Config file failed, %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unmarshal Config file failed, %v", err)
	}
}
