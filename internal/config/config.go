package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServerConfig struct {
	Host string `default:":8080"`
}

type LizaAlertAPI struct {
	EndpointURL string `required:"true"`
}

type AppConfig struct {
	HTTPServer   HTTPServerConfig
	LizaAlertAPI LizaAlertAPI
}

func ReadConfigFromEnv() AppConfig {
	var cfg AppConfig
	envconfig.MustProcess("locator", &cfg)
	envconfig.Usage("locator", &cfg)
	log.Printf("got config values: %+v", cfg)
	return cfg
}
