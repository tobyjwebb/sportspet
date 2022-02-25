package redis_test

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/teams"
	redis_team_service "github.com/tobyjwebb/teamchess/src/teams/redis"
	"github.com/tobyjwebb/teamchess/src/test"
)

func TestRedisUserService_Login(t *testing.T) {
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
