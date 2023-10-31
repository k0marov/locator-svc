package main

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/setup"
)

func main() {
	cfg := setup.ReadConfigFromEnv()
	setup.InitializeAndStart(cfg)
}
