package api

import (
	"net/http"
)

func (s Server) handleAuthentication() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
		Email  string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var request request

		err := s.decode(w, r, request)
		if err != nil {
			panic(err)
		}

		err, session := s.app.Auth(request.Email, request.Password)
		if err != nil {
			panic(err)
		}

		s.encodeAndRespond(w, r, response{
			Token:  string(session.GetToken()),
			Expire: session.GetExpire(),
			Email:  request.Email,
		}, 200)
	}
}
