package app

import (
	"account/internal/user"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (s Server) handleAuthentication() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		UserSession user.Session `json:"user_session"`
		User        user.User    `json:"user"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// todo handle it
			panic(err)
		}

		var request request
		_ = json.Unmarshal(body, request)

		w.WriteHeader(200)

		response := response{
			User:        user.NewUser(request.Email),
			UserSession: user.NewSession(),
		}

		responseJson, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(responseJson)
	}
}
