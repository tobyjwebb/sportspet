package web_frontend_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/tobyjwebb/teamchess/src/teams"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func TestTeamsHandler(t *testing.T) {
	type args struct {
		method         string
		teamName       string
		ownerSessionID string
	}
	tests := []struct {
		name        string
		args        args
		teamService teams.TeamService
		wantStatus  int
		wantTeam    *teams.Team
	}{
		{
			"Post createTeam succeeds",
			args{
				method:         http.MethodPost,
				teamName:       "theteamname",
				ownerSessionID: "ownersessionid",
			},
			&teams.TeamServiceMock{
				CreateTeamFn: func(team *teams.Team) error {
					if team.Name != "theteamname" {
						t.Errorf("Wrong team name: %q", team.Name)
					}
					team.ID = "someteamid"
					return nil
				},
			},
			http.StatusCreated,
			&teams.Team{
				ID:      "someteamid",
				Name:    "theteamname",
				Owner:   "ownersessionid",
				Members: []string{"ownersessionid"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			data.Set("name", tt.args.teamName)
			data.Set("owner", tt.args.ownerSessionID)
			request, err := http.NewRequest(tt.args.method, "dummy", strings.NewReader(data.Encode()))
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
			if err != nil {
				t.Fatalf("Got error: %v", err)
			}
			response := httptest.NewRecorder()
			server := web_frontend.NewServer(nil)
			server.TeamService = tt.teamService

			server.CreateTeamHandler(response, request)

			gotStatus := response.Result().StatusCode
			if gotStatus != tt.wantStatus {
				t.Errorf("Got status %d, want %d", gotStatus, tt.wantStatus)
			}

			var gotTeam *teams.Team
			gotBytes := response.Body.Bytes()
			if err := json.Unmarshal(gotBytes, &gotTeam); err != nil {
				t.Fatalf("Unexpected error decoding result: %v", err)
			}

			if !reflect.DeepEqual(gotTeam, tt.wantTeam) {
				t.Errorf("Got wrong team. Got:\n%v\n\nWant:\n%v\n", gotTeam, tt.wantTeam)
			}
		})
	}
}
