package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/battles"
	// "github.com/google/uuid"
	// "github.com/tobyjwebb/battlechess/src/battles"
)

const (
// battlesKey          = "battles"
// battlePropertiesKey = "battles:%s:properties"
// battleMembersKey    = "battles:%s:members"
// ownerKey            = "owner"
// nameKey             = "name"
// rankKey             = "rank"
// battleIDKey         = "battle_id"
// statusKey           = "status"
// timestampKey        = "timestamp"
)

var ctx = context.Background()

func New(client *redis.Client) (*redisBattleService, error) {
	return &redisBattleService{client}, nil
}

type redisBattleService struct {
	client *redis.Client
}

func (t *redisBattleService) Create(battle *battles.Battle) error {
	return nil // XXX implement create battle
}

func (t *redisBattleService) GetData(id string) (*battles.Battle, error) {
	return nil, nil // XXX Implement getbattledata
}

func (t *redisBattleService) Update(battle *battles.Battle) error {
	return nil // XXX implement update battle
}

// func (r *redisBattleService) CreateBattle(battle *battles.Battle) error {
// 	newBattleID := uuid.NewString()

// 	_, err := r.client.RPush(ctx, battlesKey, newBattleID).Result()
// 	if err != nil {
// 		return fmt.Errorf("could not add battle ID to battles list: %w", err)
// 	}
// 	_, err = r.client.HSet(ctx, fmt.Sprintf(battlePropertiesKey, newBattleID),
// 		nameKey, battle.Name,
// 		ownerKey, battle.Owner,
// 		rankKey, battle.Rank,
// 		battleIDKey, battle.Status.BattleID,
// 		statusKey, battle.Status.Status,
// 		timestampKey, battle.Status.Timestamp,
// 	).Result()
// 	if err != nil {
// 		return fmt.Errorf("could not set battle properties: %w", err)
// 	}
// 	for _, m := range battle.Members {
// 		_, err = r.client.RPush(ctx, fmt.Sprintf(battleMembersKey, newBattleID), m).Result()
// 		if err != nil {
// 			return fmt.Errorf("could not populate battle member list: %w", err)
// 		}
// 	}
// 	battle.ID = newBattleID
// 	// if s, err := r.sessionService.GetSession(battle.Owner); err != nil {
// 	// 	return err
// 	// } else {
// 	// 	s.BattleID = newBattleID
// 	// 	if err := r.sessionService.Update(s); err != nil {
// 	// 		return err
// 	// 	}
// 	// }

// 	return nil
// }

// func (r *redisBattleService) ListBattles() ([]battles.Battle, error) {
// 	var battlesList []battles.Battle
// 	battleIDsList, err := getAllList(ctx, battlesKey, r.client)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get battle list: %w", err)
// 	}

// 	for _, id := range battleIDsList {
// 		if battle, err := r.GetBattleData(id); err != nil {
// 			return nil, fmt.Errorf("could not obtain data for battle %s: %w", id, err)
// 		} else {
// 			battlesList = append(battlesList, *battle)
// 		}
// 	}
// 	return battlesList, nil
// }

// func (r *redisBattleService) GetBattleData(id string) (*battles.Battle, error) {
// 	t := &battles.Battle{
// 		// ID:     id,
// 		// Status: battles.BattleStatus{},
// 	}
// 	if fields, err := r.client.HGetAll(ctx, fmt.Sprintf(battlePropertiesKey, id)).Result(); err != nil {
// 		return nil, err
// 	} else {
// 		t.Name = fields[nameKey]
// 		t.Owner = fields[ownerKey]
// 		t.Rank, _ = strconv.Atoi(fields[rankKey])
// 		t.Status.Status = fields[statusKey]
// 		t.Status.BattleID = fields[battleIDKey]
// 		t.Status.Timestamp = fields[timestampKey]
// 	}

// 	if members, err := getAllList(ctx, fmt.Sprintf(battleMembersKey, id), r.client); err != nil {
// 		return nil, fmt.Errorf("could not obtain the list of battle members for %q: %w", id, err)
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

// func (r *redisBattleService) JoinBattle(sessionID, battleID string) (*battles.Battle, error) {
// 	// sess, err := r.sessionService.GetSession(sessionID)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("could not obtain session: %w", err)
// 	// }

// 	// _, err = r.client.RPush(ctx, fmt.Sprintf(battleMembersKey, battleID), sessionID).Result()
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("could not update battle member list: %w", err)
// 	// }

// 	// sess.BattleID = battleID
// 	// if err := r.sessionService.Update(sess); err != nil {
// 	// 	return nil, fmt.Errorf("could not not update session: %w", err)
// 	// }

// 	// return r.GetBattleData(battleID)
// 	return nil, nil
// }
