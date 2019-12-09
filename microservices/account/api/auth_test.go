package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3auris/auction-server/app"
)

func TestServer_RegisterAndAuthenticate(t *testing.T) {

	tests := map[string]struct {
		registerPayload    string
		registerStatusCode int

		authPayload    string
		authStatusCode int
	}{
		"auth should success": {
			registerPayload:    `{"email": "admin@admin.com","password":"admin"}`,
			registerStatusCode: http.StatusOK,

			authPayload:    `{"email": "admin@admin.com","password":"admin"}`,
			authStatusCode: http.StatusOK,
		},
		"auth should fail because user's password is wrong": {
			registerPayload:    `{"email": "admin@admin.com","password":"admin123"}`,
			registerStatusCode: http.StatusOK,

			authPayload:    `{"email": "admin@admin.com","password":"admin"}`,
			authStatusCode: http.StatusBadRequest,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			server := NewServer(app.NewApp())

			func(t *testing.T, server *Server) {
				var buf bytes.Buffer
				buf.Write([]byte(tt.registerPayload))

				r := httptest.NewRequest(http.MethodPost, apiPrefix+"/register", &buf)
				w := httptest.NewRecorder()

				server.handleRegistration()(w, r)

				if w.Code != tt.registerStatusCode {
					t.Errorf("wrong status code: got: %d, expected: %d", w.Code, tt.registerStatusCode)
				}
			}(t, server)

			func(t *testing.T, server *Server) {
				var buf bytes.Buffer
				buf.Write([]byte(tt.authPayload))

				r := httptest.NewRequest(http.MethodPost, apiPrefix+"/authenticate", &buf)
				w := httptest.NewRecorder()

				server.handleAuthentication()(w, r)

				if w.Code != tt.authStatusCode {
					t.Errorf("wrong status code: got: %d, expected: %d", w.Code, tt.authStatusCode)
				}
			}(t, server)
		})
	}
}
