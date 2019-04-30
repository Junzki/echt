package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	Bind string `env:"BIND" envDefault:"127.0.0.1"`
	Port int    `env:"PORT" envDefault:"8000"`
}

func (self *Config) LoadEnv() {
	err := env.Parse(self)
	if nil != err {
		log.Fatal("Cannot load config from env: ", err)
		os.Exit(1)
	}
}

var settings = &Config{}

func GetSettings() *Config {
	return settings
}
