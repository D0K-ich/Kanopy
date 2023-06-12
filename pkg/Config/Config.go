package config

import (
	lg "TimeTracker/pkg/Loging"
	"log"
	"sync"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	isDevelopment bool `env:"isDev" env-default:"true"` 
	isDebug bool `env:"isDebug" env-default:"false"` 
	Listen struct {
		Type string `env:"LISTEN-TYPE" env-default:"port"`
		BindIP string `env:"BIND-IP" env-default:"0.0.0.0"`
		Port string `env:"PORT" env-default:"8080"`
	}
	AppConfig struct {
		LogLevel string `env:"LOG-LEVEL" env-default:"Info"`
	}
	Width int `env:"WIDTH" env-default:"800"`
	Height int `env:"HEIGHT" env-default:"800"`
}

var instance *Config
var once sync.Once


func GetConfig() (*Config) {
	once.Do(func() {
		logger := lg.GetLogger()
		logger.Info("Start config")
		

		instance = &Config{}

		cleanenv.ReadEnv(instance)

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "There are some porblems with configuring"
			cleanenv.GetDescription(instance, &helpText)
			log.Fatal(err)
		}
		
	})
	return instance
}