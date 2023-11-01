package delivery

import (
	"encoding/json"
	"fmt"
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"net/http"
	"net/url"
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
	latitude, err := parseFloatQuery(urlQuery, "latitude")
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	longitude, err := parseFloatQuery(urlQuery, "longitude")
	if err != nil {
		core.WriteErrorResponse(w, err)
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
	json.NewEncoder(w).Encode(missing)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func parseFloatQuery(query url.Values, name string) (float64, error) {
	value, err := strconv.ParseFloat(query.Get(name), 64)
	if err != nil {
		return 0, &core.ClientError{
			DisplayMessage: fmt.Sprintf("couldn't parse '%s' url parameter as float", name),
			HTTPCode:       http.StatusBadRequest,
		}
	}
	return value, nil
}
