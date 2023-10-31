package internal

import (
	"net/http"
)
import "github.com/go-chi/chi/v5"

type ILocatorService interface {
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
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
