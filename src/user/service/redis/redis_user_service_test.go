package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	redis_user_service "github.com/tobyjwebb/teamchess/src/user/service/redis"
)

func TestRedisUserService_Login(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := setupRedis(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)

	r, err := redis_user_service.New(redisContainer.Addr)
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

	// Second login should return same session ID
	gotUser1Session2, gotErr := r.Login("user1")

	if gotUser1Session2 != gotUser1Session {
		t.Errorf("Was expecting same session ID, got different one: %s != %s", gotUser1Session, gotUser1Session2)
	}
	if gotErr != nil {
		t.Errorf("Got unexpected error: %v", gotErr)
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

type redisContainer struct {
	testcontainers.Container
	Addr string
}

func setupRedis(ctx context.Context) (*redisContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "eqalpha/keydb:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Thread 0 alive"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", hostIP, mappedPort.Port())

	return &redisContainer{Container: container, Addr: addr}, nil
}
