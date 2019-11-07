package main

import (
	"account"
	"account/pkg/token"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	prefix := "/v1/account"

	router := mux.NewRouter()

	api := account.Server{}

	router.HandleFunc(prefix+"/authenticate", api.HandleAuthentication(token.Token{})).Methods(http.MethodPost)
	router.HandleFunc(prefix+"/health", alive).Methods(http.MethodGet)
	router.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("account microservice"))
	}).Methods(http.MethodGet)

	err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router))

	if err != nil {
		fmt.Print(err)
	}
}

func alive(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Printf("could not get hostname: %v", err)
	}

	w.WriteHeader(200)
	_, _ = w.Write([]byte("alive, hostname: " + hostname))
}
