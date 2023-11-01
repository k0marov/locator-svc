package delivery

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
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
	callerLocation, err := DecodeLocationRequest(r)
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}

	missing, err := s.svc.GetRelevantMissing(*callerLocation)
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	EncodeMissingPeopleResponse(w, missing)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
