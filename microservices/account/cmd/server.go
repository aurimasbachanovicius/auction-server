package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/3auris/auction-server/app"
	"github.com/gorilla/handlers"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	server := app.NewServer()

	return http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(server))
}
