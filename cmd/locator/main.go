package main

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/config"
	"gitlab.com/samkomarov/locator-svc.git/internal/setup"
)

func main() {
	cfg := config.ReadConfigFromEnv()
	setup.InitializeAndStart(cfg)
}
