package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var ctx = context.Background()

func New(addr string) (*redisUserService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	res := rdb.Ping(ctx)
	if err := res.Err(); err != nil {
		return nil, err
	}
	return &redisUserService{client: rdb}, nil
}

type redisUserService struct {
	client *redis.Client
}

func (r *redisUserService) Login(nick string) (sessionID string, err error) {
	// XXX implement

	sessionID = uuid.NewString()
	err = fmt.Errorf("Not implemented yet!")
	return
}
