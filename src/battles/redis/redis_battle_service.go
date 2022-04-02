package redis_battle_service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/tobyjwebb/teamchess/src/battles"
)

const (
	battlesKey          = "battles"
	battlePropertiesKey = "battles:%s:properties"
	boardStatusKey      = "board_status"
	whiteTeamKey        = "white_team"
	blackTeamKey        = "black_team"
	moveCountKey        = "move_count"
)

var ctx = context.Background()

func New(client *redis.Client) (*redisBattleService, error) {
	return &redisBattleService{client}, nil
}

type redisBattleService struct {
	client *redis.Client
}

func (r *redisBattleService) Create(battle *battles.Battle) error {
	newBattleID := uuid.NewString()

	_, err := r.client.RPush(ctx, battlesKey, newBattleID).Result()
	if err != nil {
		return fmt.Errorf("could not add battle ID to battles list: %w", err)
	}
	_, err = r.client.HSet(ctx, fmt.Sprintf(battlePropertiesKey, newBattleID),
		boardStatusKey, battle.Board,
		whiteTeamKey, battle.WhiteTeamID,
		blackTeamKey, battle.BlackTeamID,
		moveCountKey, battle.MoveCount,
	).Result()
	if err != nil {
		return fmt.Errorf("could not set battle properties: %w", err)
	}
	battle.ID = newBattleID
	return nil
}

func (r *redisBattleService) GetData(id string) (*battles.Battle, error) {
	b := &battles.Battle{
		ID: id,
	}
	if fields, err := r.client.HGetAll(ctx, fmt.Sprintf(battlePropertiesKey, id)).Result(); err != nil {
		return nil, err
	} else {
		b.Board = fields[boardStatusKey]
		b.WhiteTeamID = fields[whiteTeamKey]
		b.BlackTeamID = fields[blackTeamKey]
		b.MoveCount, _ = strconv.Atoi(fields[moveCountKey])
	}

	return b, nil
}

func (r *redisBattleService) Update(battle *battles.Battle) error {
	// TODO: DRY (similar to Create())
	_, err := r.client.HSet(ctx, fmt.Sprintf(battlePropertiesKey, battle.ID),
		boardStatusKey, battle.Board,
		whiteTeamKey, battle.WhiteTeamID,
		blackTeamKey, battle.BlackTeamID,
		moveCountKey, battle.MoveCount,
	).Result()
	if err != nil {
		return fmt.Errorf("could not set battle properties: %w", err)
	}
	return nil
}
