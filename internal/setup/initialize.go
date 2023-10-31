package setup

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/delivery"
	"gitlab.com/samkomarov/locator-svc.git/internal/repository"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	repo := repository.NewExternalAPILocatorRepo(cfg.LocatorAPI)
	svc := service.NewLocatorService(repo)
	srv := delivery.NewServer(svc)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
