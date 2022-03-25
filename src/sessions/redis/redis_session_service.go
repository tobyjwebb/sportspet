package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/sessions"
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

func (s *redisSessionService) GetSession(id string) (*sessions.Session, error) {
	return nil, fmt.Errorf("GetSession: not implemented!")
}

func (s *redisSessionService) Update(session *sessions.Session) error {
	return fmt.Errorf("Update: not implemented!")
}
