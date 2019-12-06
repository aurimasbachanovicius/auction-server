package api

import (
	"net/http"
)

func (s Server) handleRegistration() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		s.decode(w, r, &req)

		if err := s.app.NewUser(req.Email, req.Password); err != nil {
			s.encodeAndRespond(w, r, errResponse{Error: err.Error()}, http.StatusBadRequest)
		}

		s.encodeAndRespond(w, r, struct{}{}, http.StatusOK)
	}
}

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
		s.decode(w, r, &request)

		session, err := s.app.Auth(request.Email, request.Password)
		if err != nil {
			s.encodeAndRespond(w, r, struct {
				Error string `json:"error"`
			}{
				Error: "wrong credentials",
			}, http.StatusBadRequest)
			return
		}

		s.encodeAndRespond(w, r, response{
			Token:  string(session.GetToken()),
			Expire: session.GetExpire(),
			Email:  request.Email,
		}, http.StatusOK)
	}
}
