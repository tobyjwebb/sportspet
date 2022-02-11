package redis_test

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	redis_user_service "github.com/tobyjwebb/teamchess/src/user/service/redis"
)

func TestRedisUserService_Login(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := redis_user_service.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_user_service.New(client)
	if err != nil {
		t.Fatalf("Could not get Redis User Service: %v", err)
	}

	// First login should return a unique session ID
	gotUser1Session, gotErr := r.Login("user1")

	if gotUser1Session == "" {
		t.Errorf("Got empty session ID")
	}
	if gotErr != nil {
		t.Errorf("Got unexpected error: %v", gotErr)
	}

	// Second login should return an error for the same user
	_, gotErr = r.Login("user1")

	if gotErr == nil {
		t.Errorf("Was expecting error, got none")
	}

	// Login with a different user should yield a different session ID
	gotUser2Session, gotErr := r.Login("user2")

	if gotUser2Session == "" {
		t.Errorf("Got empty session ID")
	}
	if gotUser2Session == gotUser1Session {
		t.Errorf("Was expecting different session ID, got same one: %s == %s", gotUser1Session, gotUser2Session)
	}
	if gotErr != nil {
		t.Errorf("Got unexpected error: %v", gotErr)
	}
}
