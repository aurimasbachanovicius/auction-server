package api

import (
	"net/http"
)

const apiPrefix string = "/v1/account"

func (s *Server) routes() {
	s.router.HandleFunc(apiPrefix+"/authenticate", s.handleAuthentication()).Methods(http.MethodPost)
	s.router.HandleFunc(apiPrefix+"/register", s.handleRegistration()).Methods(http.MethodPost)
	s.router.HandleFunc(apiPrefix+"/health", alive).Methods(http.MethodGet)
}
