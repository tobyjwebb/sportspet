package web_frontend_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	user_service "github.com/tobyjwebb/teamchess/src/user/service"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func TestLoginHandler(t *testing.T) {
	type args struct {
		method string
		user   string
	}
	tests := []struct {
		name         string
		args         args
		userService  user_service.UserService
		wantStatus   int
		wantLocation string
	}{
		{
			"Post login returns Redirect with session returned from UserService",
			args{
				method: http.MethodPost,
				user:   "myname",
			},
			&user_service.UserServiceMock{LoginFn: func(nick string) (sessionID string, err error) {
				if nick != "myname" {
					t.Errorf("Unexpected name received: %s", nick)
				}
				return "the_session_id", nil
			}},
			http.StatusTemporaryRedirect,
			"/web/session/the_session_id",
		},
		{
			"Post login without a user redirects to nick-required page",
			args{
				method: http.MethodPost,
			},
			nil,
			http.StatusTemporaryRedirect,
			"/nick-required.html",
		},
		{
			"Any method other than POST returns method not allowed",
			args{
				method: http.MethodPut,
			},
			nil,
			http.StatusMethodNotAllowed,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			data.Set("nick", tt.args.user)
			request, err := http.NewRequest(tt.args.method, "dummy", strings.NewReader(data.Encode()))
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			response := httptest.NewRecorder()
			server := web_frontend.NewServer(nil)
			server.UserService = tt.userService

			server.LoginHandler(response, request)

			gotStatus := response.Result().StatusCode
			if gotStatus != tt.wantStatus {
				t.Errorf("Got status %d, want %d", gotStatus, tt.wantStatus)
			}

			gotLocation := response.Result().Header.Get("Location")
			if gotLocation != tt.wantLocation {
				t.Errorf("Got location %s, want %s", gotLocation, tt.wantLocation)
			}
		})
	}
}
