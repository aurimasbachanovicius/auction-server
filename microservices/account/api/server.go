package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/3auris/auction-server/app"
	"github.com/gorilla/mux"
)

// Server for providing api
type Server struct {
	router *mux.Router

	app app.App
}

// NewServer creates new Server with routes and server dependencies
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

	marshaled, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
	}

	_, _ = w.Write(marshaled)
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, response interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.encodeAndRespond(w, r, struct{}{}, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		s.encodeAndRespond(w, r, struct{}{}, http.StatusBadRequest)
	}
}
