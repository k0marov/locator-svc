package delivery

import (
	"encoding/json"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"gitlab.com/samkomarov/locator-svc.git/internal/setup"
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
	missing, err := s.svc.GetRelevantMissing(service.GeoPoint{})
	if err != nil {
		setup.WriteErrorResponse(w, err)
		return
	}
	err = json.NewEncoder(w).Encode(missing)
	if err != nil {
		setup.WriteErrorResponse(w, &setup.ClientError{
			DisplayMessage: err.Error(),
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
