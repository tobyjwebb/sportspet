package redis

// import (
// 	"context"
// 	"fmt"
// 	"strconv"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/google/uuid"
// 	"github.com/tobyjwebb/teamchess/src/challenges"
// )

// const (
// 	challengesKey          = "challenges"
// 	challengePropertiesKey = "challenges:%s:properties"
// 	challengeMembersKey    = "challenges:%s:members"
// 	ownerKey          = "owner"
// 	nameKey           = "name"
// 	rankKey           = "rank"
// 	battleIDKey       = "battle_id"
// 	statusKey         = "status"
// 	timestampKey      = "timestamp"
// )

// var ctx = context.Background()

// func New(client *redis.Client) (*redisChallengeService, error) {
// 	return &redisChallengeService{client: client}, nil
// }

// type redisChallengeService struct {
// 	client *redis.Client
// }

// func (r *redisChallengeService) CreateChallenge(challenge *challenges.Challenge) error {
// 	newChallengeID := uuid.NewString()

// 	_, err := r.client.RPush(ctx, challengesKey, newChallengeID).Result()
// 	if err != nil {
// 		return fmt.Errorf("could not add challenge ID to challenges list: %w", err)
// 	}
// 	_, err = r.client.HSet(ctx, fmt.Sprintf(challengePropertiesKey, newChallengeID),
// 		nameKey, challenge.Name,
// 		ownerKey, challenge.Owner,
// 		rankKey, challenge.Rank,
// 		battleIDKey, challenge.Status.BattleID,
// 		statusKey, challenge.Status.Status,
// 		timestampKey, challenge.Status.Timestamp,
// 	).Result()
// 	if err != nil {
// 		return fmt.Errorf("could not set challenge properties: %w", err)
// 	}
// 	for _, m := range challenge.Members {
// 		_, err = r.client.RPush(ctx, fmt.Sprintf(challengeMembersKey, newChallengeID), m).Result()
// 		if err != nil {
// 			return fmt.Errorf("could not populate challenge member list: %w", err)
// 		}
// 	}
// 	challenge.ID = newChallengeID
// 	return nil
// }

// func (r *redisChallengeService) ListChallenges() ([]challenges.Challenge, error) {
// 	var challengesList []challenges.Challenge
// 	challengeIDsList, err := getAllList(ctx, challengesKey, r.client)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get challenge list: %w", err)
// 	}

// 	for _, id := range challengeIDsList {
// 		if challenge, err := r.getChallengeData(id); err != nil {
// 			return nil, fmt.Errorf("could not obtain data for challenge %s: %w", id, err)
// 		} else {
// 			challengesList = append(challengesList, *challenge)
// 		}
// 	}
// 	return challengesList, nil
// }

// func (r *redisChallengeService) getChallengeData(id string) (*challenges.Challenge, error) {
// 	t := &challenges.Challenge{
// 		ID:     id,
// 		Status: challenges.ChallengeStatus{},
// 	}
// 	if fields, err := r.client.HGetAll(ctx, fmt.Sprintf(challengePropertiesKey, id)).Result(); err != nil {
// 		return nil, err
// 	} else {
// 		t.Name = fields[nameKey]
// 		t.Owner = fields[ownerKey]
// 		t.Rank, _ = strconv.Atoi(fields[rankKey])
// 		t.Status.Status = fields[statusKey]
// 		t.Status.BattleID = fields[battleIDKey]
// 		t.Status.Timestamp = fields[timestampKey]
// 	}

// 	if members, err := getAllList(ctx, fmt.Sprintf(challengeMembersKey, id), r.client); err != nil {
// 		return nil, fmt.Errorf("could not obtain the list of challenge members for %q: %w", id, err)
// 	} else {
// 		t.Members = members
// 	}

// 	return t, nil
// }

// func getAllList(ctx context.Context, key string, r *redis.Client) ([]string, error) {
// 	count, err := r.LLen(ctx, key).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get number of items for key %q: %w", key, err)
// 	}
// 	list, err := r.LRange(ctx, key, 0, count).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get list of items for key %q: %w", key, err)
// 	}
// 	return list, nil
// }

// func (r *redisChallengeService) JoinChallenge(sessionID, challengeID string) (*challenges.Challenge, error) {
// 	return nil, fmt.Errorf("not implemented")
// }
