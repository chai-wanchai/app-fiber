package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTP        httpConfig
	SQLMasterDB sql
}

type sql struct {
	Username string `envconfig:"MYSQL_USERNAME"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	Host     string `envconfig:"MYSQL_HOST"`
	Port     string `envconfig:"MYSQL_PORT"`
	Db       string `envconfig:"MYSQL_DATABASE"`
}

type httpConfig struct {
	Host    string `envconfig:"HTTP_HOST" default:"0.0.0.0"`
	Port    string `envconfig:"APP_HTTP_PORT" default:"3002"`
	Timeout int    `envconfig:"HTTP_TIMEOUT" default:"10"`
}

var cfg Config

func New() {
	_ = godotenv.Load()
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("read env error %v", err)
	}
}

func GetConfig() Config {
	return cfg
}
