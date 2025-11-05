package config

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	Port       uint16 `env:"PORT" envDefault:"3000"`
	StreamsURL string `env:"STREAMS_URL" envDefault:"https://docs.google.com/spreadsheets/d/e/2PACX-1vQcJWZc9LXpAjKWbMNFrI7Gbry3GfkfsK55k8Mp3EW65rIJhAZZG0W9WGwgrSwAB5J8iaVTeKFWh2Or/pub?output=csv"`
}

var envConfig *Env

func GetEnv() *Env {
	if envConfig == nil {
		cfg := &Env{}
		if err := env.Parse(cfg); err != nil {
			panic(err)
		}
		envConfig = cfg
	}
	return envConfig
}
