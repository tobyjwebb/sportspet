package web_frontend_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tobyjwebb/teamchess/src/teams"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func TestCreateTeamsHandler(t *testing.T) {
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

func TestListTeamsHandlerSuccess(t *testing.T) {
	someTeam := teams.Team{
		ID:      "someid",
		Name:    "someName",
		Owner:   "someowner",
		Rank:    99,
		Members: []string{"mem1", "mem2"},
		Status: teams.TeamStatus{
			BattleID:  "mybattle",
			Status:    "myfoostatus",
			Timestamp: "somets",
		},
	}
	request, err := http.NewRequest(http.MethodGet, "dummy", nil)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	response := httptest.NewRecorder()
	server := web_frontend.NewServer(nil)
	server.TeamService = &teams.TeamServiceMock{
		ListTeamsFn: func() ([]teams.Team, error) {
			return []teams.Team{someTeam}, nil
		},
	}

	server.ListTeamsHandler(response, request)

	gotStatus := response.Result().StatusCode
	wantStatus := http.StatusOK
	if gotStatus != wantStatus {
		t.Errorf("Got status %d, want %d", gotStatus, wantStatus)
	}

	var gotTeams []teams.Team
	gotBytes := response.Body.Bytes()
	if err := json.Unmarshal(gotBytes, &gotTeams); err != nil {
		t.Fatalf("Unexpected error decoding result: %v", err)
	}

	wantTeams := []teams.Team{someTeam}
	if !reflect.DeepEqual(gotTeams, wantTeams) {
		t.Errorf("Got wrong team list. Got:\n%v\n\nWant:\n%v\n", gotTeams, wantTeams)
	}
}

func TestListTeamsHandlerError(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "dummy", nil)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	response := httptest.NewRecorder()
	server := web_frontend.NewServer(nil)
	server.TeamService = &teams.TeamServiceMock{
		ListTeamsFn: func() ([]teams.Team, error) {
			return nil, fmt.Errorf("OhNoSomethingBadHappened")
		},
	}

	server.ListTeamsHandler(response, request)

	gotStatus := response.Result().StatusCode
	wantStatus := http.StatusInternalServerError
	if gotStatus != wantStatus {
		t.Errorf("Got status %d, want %d", gotStatus, wantStatus)
	}
}

// func TestJoinTeamsHandlerError(t *testing.T) {
// }

func TestSpec(t *testing.T) {

	Convey("Given a server", t, func() {
		server := web_frontend.NewServer(nil)

		Convey("Given a request to join a team", func() {
			sessionID := "mysessionid"
			teamID := "foo-team-id"
			url := fmt.Sprintf("/api/v1/teams/%s/join", teamID)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			if err != nil {
				t.Fatalf("Got error: %v", err)
			}
			request.Header.Set("Authorization", "Bearer "+sessionID)
			response := httptest.NewRecorder()

			Convey("Given a teamservice mock that returns an error", func() {
				var gotSessionID string
				var gotTeamID string
				server.TeamService = &teams.TeamServiceMock{
					JoinTeamFn: func(sessionID, teamID string) (*teams.Team, error) {
						gotSessionID, gotTeamID = sessionID, teamID
						return nil, fmt.Errorf("ohnosomethingbadhappened")
					},
				}

				Convey("When the handler is called", func() {
					server.ServeHTTP(response, request)

					Convey("Then the status code should be internal server error", func() {
						So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
					})

					Convey("Then arguments to join are the ones set in the request", func() {
						So(gotSessionID, ShouldEqual, sessionID)
						So(gotTeamID, ShouldEqual, teamID)
					})
				})
			})
		})
	})
}
