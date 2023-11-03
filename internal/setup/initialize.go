package setup

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/config"
	"gitlab.com/samkomarov/locator-svc.git/internal/delivery"
	"gitlab.com/samkomarov/locator-svc.git/internal/repository"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"log"
	"net/http"
)

func InitializeAndStart(cfg config.AppConfig) {
	repo := repository.NewExternalAPILocatorRepo(cfg.LizaAlertAPI)
	svc := service.NewLocatorService(repo)
	srv := delivery.NewServer(svc)
	log.Printf("Listening at %s", cfg.HTTPServer.Host)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
