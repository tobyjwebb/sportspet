package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const sessionsKey = "sessions"

var ctx = context.Background()

func New(client *redis.Client) (*redisSessionService, error) {

	return &redisSessionService{client: client}, nil
}

type redisSessionService struct {
	client *redis.Client
}

func (r *redisSessionService) Login(nick string) (sessionID string, err error) {
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
		return "", nil
	}
	return
}
