package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/sessions"
	"github.com/tobyjwebb/teamchess/src/teams"
)

const (
	teamsKey          = "teams"
	teamPropertiesKey = "teams:%s:properties"
	teamMembersKey    = "teams:%s:members"
	ownerKey          = "owner"
	nameKey           = "name"
	rankKey           = "rank"
	battleIDKey       = "battle_id"
	statusKey         = "status"
	timestampKey      = "timestamp"
)

var ctx = context.Background()

func New(client *redis.Client, sessionService sessions.SessionService) (*redisTeamService, error) {
	return &redisTeamService{client: client, sessionService: sessionService}, nil
}

type redisTeamService struct {
	client         *redis.Client
	sessionService sessions.SessionService
}

func (r *redisTeamService) CreateTeam(team *teams.Team) error {
	newTeamID := uuid.NewString()

	_, err := r.client.RPush(ctx, teamsKey, newTeamID).Result()
	if err != nil {
		return fmt.Errorf("could not add team ID to teams list: %w", err)
	}
	_, err = r.client.HSet(ctx, fmt.Sprintf(teamPropertiesKey, newTeamID),
		nameKey, team.Name,
		ownerKey, team.Owner,
		rankKey, team.Rank,
		battleIDKey, team.Status.BattleID,
		statusKey, team.Status.Status,
		timestampKey, team.Status.Timestamp,
	).Result()
	if err != nil {
		return fmt.Errorf("could not set team properties: %w", err)
	}
	for _, m := range team.Members {
		_, err = r.client.RPush(ctx, fmt.Sprintf(teamMembersKey, newTeamID), m).Result()
		if err != nil {
			return fmt.Errorf("could not populate team member list: %w", err)
		}
	}
	team.ID = newTeamID
	if s, err := r.sessionService.GetSession(team.Owner); err != nil {
		return err
	} else {
		s.TeamID = newTeamID
		if err := r.sessionService.Update(s); err != nil {
			return err
		}
	}

	return nil
}

// TODO: DRY Update Team
func (r *redisTeamService) Update(team *teams.Team) error {
	_, err := r.client.HSet(ctx, fmt.Sprintf(teamPropertiesKey, team.ID),
		nameKey, team.Name,
		ownerKey, team.Owner,
		rankKey, team.Rank,
		battleIDKey, team.Status.BattleID,
		statusKey, team.Status.Status,
		timestampKey, team.Status.Timestamp,
	).Result()
	return err
}

func (r *redisTeamService) ListTeams() ([]teams.Team, error) {
	var teamsList []teams.Team
	teamIDsList, err := getAllList(ctx, teamsKey, r.client)
	if err != nil {
		return nil, fmt.Errorf("could not get team list: %w", err)
	}

	for _, id := range teamIDsList {
		if team, err := r.GetTeamData(id); err != nil {
			return nil, fmt.Errorf("could not obtain data for team %s: %w", id, err)
		} else {
			teamsList = append(teamsList, *team)
		}
	}
	return teamsList, nil
}

func (r *redisTeamService) GetTeamData(id string) (*teams.Team, error) {
	t := &teams.Team{
		ID:     id,
		Status: teams.TeamStatus{},
	}
	if fields, err := r.client.HGetAll(ctx, fmt.Sprintf(teamPropertiesKey, id)).Result(); err != nil {
		return nil, err
	} else {
		t.Name = fields[nameKey]
		t.Owner = fields[ownerKey]
		t.Rank, _ = strconv.Atoi(fields[rankKey])
		t.Status.Status = fields[statusKey]
		t.Status.BattleID = fields[battleIDKey]
		t.Status.Timestamp = fields[timestampKey]
	}

	if members, err := getAllList(ctx, fmt.Sprintf(teamMembersKey, id), r.client); err != nil {
		return nil, fmt.Errorf("could not obtain the list of team members for %q: %w", id, err)
	} else {
		t.Members = members
	}

	return t, nil
}

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

func (r *redisTeamService) JoinTeam(sessionID, teamID string) (*teams.Team, error) {
	sess, err := r.sessionService.GetSession(sessionID)
	if err != nil {
		return nil, fmt.Errorf("could not obtain session: %w", err)
	}

	_, err = r.client.RPush(ctx, fmt.Sprintf(teamMembersKey, teamID), sessionID).Result()
	if err != nil {
		return nil, fmt.Errorf("could not update team member list: %w", err)
	}

	sess.TeamID = teamID
	if err := r.sessionService.Update(sess); err != nil {
		return nil, fmt.Errorf("could not not update session: %w", err)
	}

	return r.GetTeamData(teamID)
}
