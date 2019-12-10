package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/3auris/auction-server/api"
	"github.com/3auris/auction-server/app"
	"github.com/gorilla/handlers"
)

func main() {
	if err := run(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}

func run() error {
	server := api.NewServer(app.NewApp())

	return http.ListenAndServe(
		":3000",
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(server),
	)
}
