package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const sessionsKey = "sessions"

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
	res, err := r.client.HGet(ctx, sessionsKey, nick).Result()
	if err == redis.Nil {
		// Create and store new sessionid
		sessionID = uuid.NewString()
		_, err = r.client.HSet(ctx, sessionsKey, nick, sessionID).Result()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		sessionID = res
	}
	return
}
