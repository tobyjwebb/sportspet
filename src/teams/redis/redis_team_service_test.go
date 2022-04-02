package redis_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tobyjwebb/teamchess/src/sessions"
	"github.com/tobyjwebb/teamchess/src/teams"
	redis_team_service "github.com/tobyjwebb/teamchess/src/teams/redis"
	"github.com/tobyjwebb/teamchess/src/test"
)

var ctx = context.Background()

func TestRedisTeamsService_CreateTeam(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	mockSessionService := &sessions.SessionServiceMock{}
	var gotGetSessionIDParam, gotUpdateSessionNewTeamID string
	mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
		gotGetSessionIDParam = id
		return &sessions.Session{}, nil
	}
	mockSessionService.UpdateFn = func(s *sessions.Session) error {
		gotUpdateSessionNewTeamID = s.TeamID
		return nil
	}

	// TODO: DRY here:
	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_team_service.New(client, mockSessionService)
	if err != nil {
		t.Fatalf("Could not get Redis Team Service: %v", err)
	}

	sessionID := "test-session-id"
	team := &teams.Team{Owner: sessionID}
	gotErr := r.CreateTeam(team)

	if gotErr != nil {
		t.Errorf("Got unexpected error: %v", gotErr)
	}

	if gotGetSessionIDParam != sessionID {
		t.Errorf("got wrong session ID param: %q, want %q", gotGetSessionIDParam, sessionID)
	}

	if team.ID == "" {
		t.Errorf("Team ID has not been initialized")
	}

	if gotUpdateSessionNewTeamID != team.ID {
		t.Errorf("got wrong new team ID in UpdateSession: %q, want %q", gotUpdateSessionNewTeamID, team.ID)
	}
}

func TestRedisTeamsService_ListTeams(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	mockSessionService := &sessions.SessionServiceMock{}
	mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
		return &sessions.Session{}, nil
	}
	mockSessionService.UpdateFn = func(s *sessions.Session) error {
		return nil
	}

	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_team_service.New(client, mockSessionService)
	if err != nil {
		t.Fatalf("Could not get Redis Team Service: %v", err)
	}

	team := &teams.Team{
		Name:  "somename",
		Owner: "someowner",
		Rank:  9,
		Status: teams.TeamStatus{
			BattleID:  "mybattleid",
			Status:    "hello world",
			Timestamp: "testtimestamp",
		},
		Members: []string{"foo", "bar"},
	}
	_ = r.CreateTeam(team)

	gotTeamList, err := r.ListTeams()
	if err != nil {
		t.Fatalf("Could not get team list: %v", err)
	}

	wantTeamList := []teams.Team{*team}

	if !reflect.DeepEqual(gotTeamList, wantTeamList) {
		t.Errorf("Got wrong team list. Got:\n%+v\n\nWant:\n%+v\n", gotTeamList, wantTeamList)
	}
}

func TestTeamService_JoinTeam(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	Convey("Given a teamservice", t, func() {
		mockSessionService := &sessions.SessionServiceMock{}
		r, err := redis_team_service.New(client, mockSessionService)
		if err != nil {
			t.Fatalf("Could not get Redis Team Service: %v", err)
		}

		Convey("Given a session that does not exist", func() {
			sessionID := "dummy-404-session"
			teamID := "dummy-team"

			var gotSessionIDParam string
			mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
				gotSessionIDParam = id
				return nil, fmt.Errorf("could not find the session")
			}

			Convey("When JoinTeam is called", func() {
				_, err := r.JoinTeam(sessionID, teamID)
				So(gotSessionIDParam, ShouldEqual, sessionID)

				Convey("Then an error is returned", func() {
					So(err, ShouldBeError)
				})
			})
		})

		Convey("Given a session that exists", func() {
			sessionID := "a-session-id-that-exists"
			teamID := "some-team"
			var gotNewTeamID string
			mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
				return &sessions.Session{ID: sessionID}, nil
			}
			mockSessionService.UpdateFn = func(s *sessions.Session) error {
				gotNewTeamID = s.TeamID
				return nil
			}
			testTeam := &teams.Team{
				Owner: "some-team-owner",
			}
			err := r.CreateTeam(testTeam)
			So(err, ShouldBeNil)

			Convey("When JoinTeam is called", func() {
				updatedTeamData, err := r.JoinTeam(sessionID, teamID)

				Convey("There is no error returned", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then the session is added to the team members", func() {
					So(updatedTeamData, ShouldNotBeNil)
					So(sessionID, ShouldBeIn, updatedTeamData.Members)
				})

				Convey("Then the session's team is updated to the new team", func() {
					So(gotNewTeamID, ShouldEqual, teamID)
				})
			})
		})
	})
}
