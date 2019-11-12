package app

import (
	"net/http"
)

func (s *Server) routes() {
	prefix := "/v1/account"

	s.router.HandleFunc(prefix+"/authenticate", s.handleAuthentication()).Methods(http.MethodPost)
	s.router.HandleFunc(prefix+"/health", alive).Methods(http.MethodGet)
}
