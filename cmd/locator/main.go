package main

import "gitlab.com/samkomarov/locator-svc.git/internal"

func main() {
	cfg := internal.ReadConfigFromEnv()
	internal.InitializeAndStart(cfg)
}
