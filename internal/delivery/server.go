package delivery

import (
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "gitlab.com/samkomarov/locator-svc.git/docs"
	"gitlab.com/samkomarov/locator-svc.git/internal/config"
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"net/http"
)
import "github.com/go-chi/chi/v5"

type ILocatorService interface {
	GetRelevantMissing(aroundPoint service.GeoPoint) ([]service.MissingPerson, error)
}

type Server struct {
	cfg config.HTTPServerConfig
	svc ILocatorService
	r   chi.Router
}

func NewServer(cfg config.HTTPServerConfig, svc ILocatorService) http.Handler {
	srv := &Server{cfg, svc, chi.NewRouter()}
	srv.defineEndpoints()
	return srv
}

//	@title			Locator API
//	@version		1.0
//	@description	An API for getting the most relevant missing people info from LizaAlerts.

//	@contact.name	Sam Komarov
//	@contact.url	github.com/k0marov
//	@contact.email	sam@skomarov.com

// @host		sber.skomarov.com
// @BasePath	/api/v1
// @schemes     https http
func (s *Server) defineEndpoints() {
	s.r.Get("/swagger/*", httpSwagger.WrapHandler)
	s.r.Route("/api/v1/locator", func(r chi.Router) {
		r.Get("/missing", s.GetMissing)
	})
}

// GetMissing godoc
//
//		@Summary		Get info about relevant missing people
//		@Description	Given a geo location, returns a list of people from LizaAlertAPI,
//	    @Description 	ordered by their relevance (calculated from distance and days missing) to provided location.
//		@Tags			locator
//		@Produce		json
//		@Param			latitude	query		float64	true	"Geo location latitude"
//		@Param			longitude	query		float64	true	"Geo location longitude"
//		@Success		200	{object}	[]MissingPersonResponse
//		@Router			/locator/missing [get]
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
