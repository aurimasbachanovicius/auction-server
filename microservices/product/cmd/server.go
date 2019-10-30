package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	prefix := "/v1/product"

	router := mux.NewRouter()

	router.HandleFunc(prefix+"/health", alive).Methods(http.MethodGet)
	router.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("product microservice"));
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
