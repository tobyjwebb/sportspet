package service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tobyjwebb/teamchess/src/user/service"
)

func TestLoginRedirectsWhenNoNick(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/login", nil)
	response := httptest.NewRecorder()

	service.LoginHandler(response, request)

	gotStatus := response.Result().StatusCode
	wantStatus := http.StatusTemporaryRedirect

	gotLocation := response.Result().Header.Get("Location")
	wantLocation := "/"

	if gotStatus != wantStatus {
		t.Errorf("Got status %d, want %d", gotStatus, wantStatus)
	}
	if gotLocation != wantLocation {
		t.Errorf("Got location %s, want %s", gotLocation, wantLocation)
	}
}

func TestLoginRedirectsWhenNickAvailable(t *testing.T) {
	nick := "nick=dummy"
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(nick))
	response := httptest.NewRecorder()

	service.LoginHandler(response, request)

	gotStatus := response.Result().StatusCode
	wantStatus := http.StatusTemporaryRedirect

	gotLocation := response.Result().Header.Get("Location")
	wantLocation := "/web/session/xxxtobeimplemented"

	if gotStatus != wantStatus {
		t.Errorf("Got status %d, want %d", gotStatus, wantStatus)
	}
	if gotLocation != wantLocation {
		t.Errorf("Got location %s, want %s", gotLocation, wantLocation)
	}
}
