package redis

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RedisTestContainer struct {
	testcontainers.Container
	Addr string
}

func SetupRedisTestContainer(ctx context.Context) (*RedisTestContainer, error) {
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

	return &RedisTestContainer{Container: container, Addr: addr}, nil
}
