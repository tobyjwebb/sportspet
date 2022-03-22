package redis_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/teams"
	redis_team_service "github.com/tobyjwebb/teamchess/src/teams/redis"
	"github.com/tobyjwebb/teamchess/src/test"
)

func TestRedisTeamsService_CreateTeam(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_team_service.New(client)
	if err != nil {
		t.Fatalf("Could not get Redis Team Service: %v", err)
	}

	team := &teams.Team{}
	gotErr := r.CreateTeam(team)

	if gotErr != nil {
		t.Errorf("Got unexpected error: %v", gotErr)
	}

	if team.ID == "" {
		t.Errorf("Team ID has not been initialized")
	}
}

func TestRedisTeamsService_ListTeams(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_team_service.New(client)
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
