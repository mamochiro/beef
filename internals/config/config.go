package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
)

type Configuration struct {
	AppName         string `env:"APP_NAME" envDefault:"worker-service"`
	Port            int    `env:"PORT" envDefault:"3001"`
	HttpPort        int    `env:"HTTP_PORT" envDefault:"3002"`
	Env             string `env:"ENV" envDefault:"localhost"`
	DbDriver        string `env:"DB_DRIVER" envDefault:"postgres"`
	DbHost          string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort          string `env:"DB_PORT" envDefault:"5432"`
	DbUser          string `env:"DB_USER" envDefault:"postgres"`
	DbName          string `env:"DB_NAME" envDefault:"common"`
	DbPassword      string `env:"DB_PASSWORD" envDefault:"q:,};+(ahY)^Z$>R"`
	JaegerAgentHost string `env:"JAEGER_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_PORT" envDefault:"6831"`
	BeefEndpoint    string `env:"BEEF_ENDPOINT" envDefault:"https://baconipsum.com/api"`
}

func NewConfiguration() Configuration {
	config := Configuration{}

	if err := env.Parse(&config); err != nil {
		fmt.Errorf("%+v\n", err)
	}

	return config
}
