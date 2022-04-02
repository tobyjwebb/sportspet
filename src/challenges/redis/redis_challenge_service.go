package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/challenges"
)

const (
	teamChallengesKey      = "challenges:team:%s"
	challengePropertiesKey = "challenges:%s:properties"
	challengerKey          = "challenger"
	challengeeKey          = "challengee"
)

var ctx = context.Background()

func New(client *redis.Client) (*redisChallengeService, error) {
	return &redisChallengeService{client: client}, nil
}

type redisChallengeService struct {
	client *redis.Client
}

func (r *redisChallengeService) Create(challenge *challenges.Challenge) error {
	newChallengeID := uuid.NewString()

	for _, v := range []string{challenge.ChallengeeTeamID, challenge.ChallengerTeamID} {
		teamChallenge := fmt.Sprintf(teamChallengesKey, v)
		_, err := r.client.RPush(ctx, teamChallenge, newChallengeID).Result()
		if err != nil {
			return fmt.Errorf("could not add challenge ID to challenges list: %w", err)
		}
	}

	if _, err := r.client.HSet(ctx, fmt.Sprintf(challengePropertiesKey, newChallengeID),
		challengerKey, challenge.ChallengerTeamID,
		challengeeKey, challenge.ChallengeeTeamID,
	).Result(); err != nil {
		return fmt.Errorf("could not set challenge properties: %w", err)
	}
	challenge.ID = newChallengeID
	return nil
}

func (r *redisChallengeService) List(teamID string) ([]challenges.Challenge, error) {
	var challengesList []challenges.Challenge
	key := fmt.Sprintf(teamChallengesKey, teamID)
	challengeIDsList, err := getAllList(ctx, key, r.client)
	if err != nil {
		return nil, fmt.Errorf("could not get challenge list: %w", err)
	}

	for _, id := range challengeIDsList {
		if challenge, err := r.getChallengeData(id); err != nil {
			return nil, fmt.Errorf("could not obtain data for challenge %s: %w", id, err)
		} else {
			challengesList = append(challengesList, *challenge)
		}
	}
	return challengesList, nil
}

func (t *redisChallengeService) Delete(challengeID string) error {
	return fmt.Errorf("not implemented") // XXX implement Challenges.Delete
}

func (r *redisChallengeService) getChallengeData(id string) (*challenges.Challenge, error) {
	t := &challenges.Challenge{
		ID: id,
	}
	if fields, err := r.client.HGetAll(ctx, fmt.Sprintf(challengePropertiesKey, id)).Result(); err != nil {
		return nil, err
	} else {
		t.ChallengerTeamID = fields[challengerKey]
		t.ChallengeeTeamID = fields[challengeeKey]
	}

	return t, nil
}

// TODO this has been copy&pasted - refactor
func getAllList(ctx context.Context, key string, r *redis.Client) ([]string, error) {
	count, err := r.LLen(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("could not get number of items for key %q: %w", key, err)
	}
	list, err := r.LRange(ctx, key, 0, count).Result()
	if err != nil {
		return nil, fmt.Errorf("could not get list of items for key %q: %w", key, err)
	}
	return list, nil
}
