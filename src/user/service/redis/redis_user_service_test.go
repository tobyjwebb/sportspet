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

	gotSession, gotErr := r.Login("dummy")

	if gotSession == "" {
		t.Errorf("Got empty session ID")
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
