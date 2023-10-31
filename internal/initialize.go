package internal

import (
	"log"
	"net/http"
)

func InitializeAndStart(cfg AppConfig) {
	repo := NewExternalAPILocatorRepo()
	svc := NewLocatorService(repo)
	srv := NewServer(svc)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
