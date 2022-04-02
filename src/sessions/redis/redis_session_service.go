package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/sessions"
)

const (
	sessionsKey           = "sessions" // map of nick->sessionID
	sessionPropertiesTplt = "sessions:%s:properties"
	teamIDKey             = "team-id"
	battleIDKey           = "battle-id"
)

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
		// Create the session properties hash:
		if err := r.Update(&sessions.Session{ID: sessionID}); err != nil {
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
	sess := &sessions.Session{ID: id}
	if fields, err := s.client.HGetAll(ctx, fmt.Sprintf(sessionPropertiesTplt, id)).Result(); err != nil {
		return nil, fmt.Errorf("could not obtain session data: %w", err)
	} else {
		sess.TeamID = fields[teamIDKey]
	}

	return sess, nil
}

func (s *redisSessionService) Update(session *sessions.Session) error {
	if _, err := s.client.HSet(
		ctx,
		fmt.Sprintf(sessionPropertiesTplt, session.ID),
		teamIDKey, session.TeamID,
	).Result(); err != nil {
		return fmt.Errorf("could not update session data: %w", err)
	}
	return nil
}
