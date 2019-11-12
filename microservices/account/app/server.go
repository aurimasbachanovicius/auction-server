package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	server := Server{
		router: mux.NewRouter(),
	}
	server.routes()

	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
