package delivery

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"net/http"
)
import "github.com/go-chi/chi/v5"

type ILocatorService interface {
	GetRelevantMissing(aroundPoint service.GeoPoint) ([]service.MissingPerson, error)
}

type Server struct {
	svc ILocatorService
	r   chi.Router
}

func NewServer(svc ILocatorService) http.Handler {
	srv := &Server{svc, chi.NewRouter()}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	s.r.Route("/api/v1/locator", func(r chi.Router) {
		r.Get("/missing", s.GetMissing)
	})
}

func (s *Server) GetMissing(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
