package setup

import (
	"github.com/kelseyhightower/envconfig"
)

type HTTPServerConfig struct {
	Host string `default:"127.0.0.1:8001"`
}

type AppConfig struct {
	HTTPServer HTTPServerConfig
}

func ReadConfigFromEnv() AppConfig {
	var cfg AppConfig
	envconfig.MustProcess("locator", &cfg)
	return cfg
}
