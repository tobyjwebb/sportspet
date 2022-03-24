package redis_test

// import (
// 	"context"
// 	"reflect"
// 	"testing"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/tobyjwebb/teamchess/src/challenges"
// 	redis_challenge_service "github.com/tobyjwebb/teamchess/src/challenges/redis"
// 	"github.com/tobyjwebb/teamchess/src/test"
// )

// func TestRedischallengesService_Createchallenge(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test")
// 	}

// 	ctx := context.Background()

// 	redisContainer, err := test.SetupRedisTestContainer(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer redisContainer.Terminate(ctx)
// 	client := redis.NewClient(&redis.Options{
// 		Addr: redisContainer.Addr,
// 	})

// 	r, err := redis_challenge_service.New(client)
// 	if err != nil {
// 		t.Fatalf("Could not get Redis challenge Service: %v", err)
// 	}

// 	challenge := &challenges.challenge{}
// 	gotErr := r.Createchallenge(challenge)

// 	if gotErr != nil {
// 		t.Errorf("Got unexpected error: %v", gotErr)
// 	}

// 	if challenge.ID == "" {
// 		t.Errorf("challenge ID has not been initialized")
// 	}
// }

// func TestRedischallengesService_Listchallenges(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test")
// 	}

// 	ctx := context.Background()

// 	redisContainer, err := test.SetupRedisTestContainer(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer redisContainer.Terminate(ctx)
// 	client := redis.NewClient(&redis.Options{
// 		Addr: redisContainer.Addr,
// 	})

// 	r, err := redis_challenge_service.New(client)
// 	if err != nil {
// 		t.Fatalf("Could not get Redis challenge Service: %v", err)
// 	}

// 	challenge := &challenges.challenge{
// 		Name:  "somename",
// 		Owner: "someowner",
// 		Rank:  9,
// 		Status: challenges.challengeStatus{
// 			BattleID:  "mybattleid",
// 			Status:    "hello world",
// 			Timestamp: "testtimestamp",
// 		},
// 		Members: []string{"foo", "bar"},
// 	}
// 	_ = r.Createchallenge(challenge)

// 	gotchallengeList, err := r.Listchallenges()
// 	if err != nil {
// 		t.Fatalf("Could not get challenge list: %v", err)
// 	}

// 	wantchallengeList := []challenges.challenge{*challenge}

// 	if !reflect.DeepEqual(gotchallengeList, wantchallengeList) {
// 		t.Errorf("Got wrong challenge list. Got:\n%+v\n\nWant:\n%+v\n", gotchallengeList, wantchallengeList)
// 	}
// }
