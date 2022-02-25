package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/teams"
)

const (
	teamsKey          = "teams"
	teamPropertiesKey = "teams:%s:properties"
	teamMembersKey    = "teams:%s:members"
)

var ctx = context.Background()

func New(client *redis.Client) (*redisTeamService, error) {
	return &redisTeamService{client: client}, nil
}

type redisTeamService struct {
	client *redis.Client
}

func (r *redisTeamService) CreateTeam(team *teams.Team) error {
	newTeamID := uuid.NewString()

	_, err := r.client.RPush(ctx, teamsKey, newTeamID).Result()
	if err != nil {
		return fmt.Errorf("could not add team ID to teams list: %w", err)
	}
	_, err = r.client.HSet(ctx, fmt.Sprintf(teamPropertiesKey, newTeamID), "name", team.Name, "owner", team.Owner).Result()
	if err != nil {
		return fmt.Errorf("could not set team properties: %w", err)
	}
	_, err = r.client.RPush(ctx, fmt.Sprintf(teamMembersKey, newTeamID), team.Owner).Result()
	if err != nil {
		return fmt.Errorf("could not add owner to team member list: %w", err)
	}
	team.ID = newTeamID
	return nil
}
