package api

import (
	"fmt"
	"net/http"
	"os"
)

func alive(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Printf("could not get hostname: %v", err)
	}

	w.WriteHeader(200)
	_, _ = w.Write([]byte("alive, hostname: " + hostname))
}
