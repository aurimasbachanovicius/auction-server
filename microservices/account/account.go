package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type StringGenerator interface {
	Generate() string
}

type Server struct {
	StringGenerator StringGenerator
}

func (s Server) HandleAuthentication() http.HandlerFunc {
	type AuthenticationRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type AuthenticationResponse struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
		User   User   `json:"user"`
		New    bool   `json:"new"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// todo handle it
			panic(err)
		}

		var request AuthenticationRequest
		_ = json.Unmarshal(body, request)

		w.WriteHeader(200)

		response := AuthenticationResponse{
			Token:  s.StringGenerator.Generate(),
			Expire: "2021-01-01",
			User:   NewUser(request.Email),
			New:    true,
		}

		responseJson, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(responseJson)
	}
}
