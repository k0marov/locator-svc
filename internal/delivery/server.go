package delivery

import (
	"encoding/json"
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"net/http"
	"strconv"
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
	urlQuery := r.URL.Query()
	latitude, err := strconv.ParseFloat(urlQuery.Get("latitude"), 64)
	if err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: "couldn't parse 'latitude' url parameter as float",
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}
	longitude, err := strconv.ParseFloat(urlQuery.Get("longitude"), 64)
	if err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: "couldn't parse 'longitude' url parameter as float",
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}

	missing, err := s.svc.GetRelevantMissing(service.GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	})
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	err = json.NewEncoder(w).Encode(missing)
	if err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: err.Error(),
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
