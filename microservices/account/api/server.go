package api

import (
	"encoding/json"
	"net/http"

	"github.com/3auris/auction-server/app"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router

	app app.App
}

func NewServer(app app.App) *Server {
	server := Server{
		router: mux.NewRouter(),
		app:    app,
	}

	server.routes()

	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (Server) encodeAndRespond(w http.ResponseWriter, r *http.Request, response interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if response != nil {
		// @todo handle error
		_ = json.NewEncoder(w).Encode(response)
	}
}

func (Server) decode(w http.ResponseWriter, r *http.Request, response interface{}) error {
	return json.NewDecoder(r.Body).Decode(response)
}
