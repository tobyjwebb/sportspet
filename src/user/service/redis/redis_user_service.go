package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const sessionsKey = "sessions"

var ctx = context.Background()

func New(client *redis.Client) (*redisUserService, error) {

	return &redisUserService{client: client}, nil
}

type redisUserService struct {
	client *redis.Client
}

func (r *redisUserService) Login(nick string) (sessionID string, err error) {
	_, err = r.client.HGet(ctx, sessionsKey, nick).Result()
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
		return "", fmt.Errorf("nick %s is already in use", nick)
	}
	return
}
