package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3auris/auction-server/app"
)

func TestServer_handleAuthentication(t *testing.T) {
	payload := []byte(`{
	"email": "admin@admin.com",
	"password":"admin"
}`)
	var buf bytes.Buffer
	buf.Write(payload)

	srv := NewServer(app.NewApp())

	r := httptest.NewRequest(http.MethodPost, apiPrefix+"/authenticate", &buf)
	w := httptest.NewRecorder()

	srv.handleAuthentication()(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status code: got: %d, expected: %d", w.Code, http.StatusOK)
	}

	type resp struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
		Email  string `json:"email"`
	}

	var response resp
	body, _ := ioutil.ReadAll(w.Result().Body)
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if response.Email != "admin@admin.com" {
		t.Errorf("wrong email: got %s, expected %s", response.Email, "admin@admin.com")
	}

	if response.Expire != "2020-01-20T14:48" {
		t.Errorf("wrong expire date: got %s, expected %s", response.Email, "2020-01-20T14:48")
	}

	if response.Token == "" {
		t.Error("wrong token, should be not be empty")
	}
}
