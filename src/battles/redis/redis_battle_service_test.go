package redis_test

// import (
// 	"context"
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	"github.com/go-redis/redis/v8"
// 	. "github.com/smartystreets/goconvey/convey"
// 	"github.com/tobyjwebb/battlechess/src/battles"
// 	redis_battle_service "github.com/tobyjwebb/battlechess/src/battles/redis"
// 	"github.com/tobyjwebb/battlechess/src/sessions"
// 	"github.com/tobyjwebb/battlechess/src/test"
// )

// var ctx = context.Background()

// func TestRedisBattlesService_CreateBattle(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test")
// 	}
// 	mockSessionService := &sessions.SessionServiceMock{}
// 	var gotGetSessionIDParam, gotUpdateSessionNewBattleID string
// 	mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
// 		gotGetSessionIDParam = id
// 		return &sessions.Session{}, nil
// 	}
// 	mockSessionService.UpdateFn = func(s *sessions.Session) error {
// 		gotUpdateSessionNewBattleID = s.BattleID
// 		return nil
// 	}

// 	// TODO: DRY here:
// 	redisContainer, err := test.SetupRedisTestContainer(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer redisContainer.Terminate(ctx)
// 	client := redis.NewClient(&redis.Options{
// 		Addr: redisContainer.Addr,
// 	})

// 	r, err := redis_battle_service.New(client, mockSessionService)
// 	if err != nil {
// 		t.Fatalf("Could not get Redis Battle Service: %v", err)
// 	}

// 	sessionID := "test-session-id"
// 	battle := &battles.Battle{Owner: sessionID}
// 	gotErr := r.CreateBattle(battle)

// 	if gotErr != nil {
// 		t.Errorf("Got unexpected error: %v", gotErr)
// 	}

// 	if gotGetSessionIDParam != sessionID {
// 		t.Errorf("got wrong session ID param: %q, want %q", gotGetSessionIDParam, sessionID)
// 	}

// 	if battle.ID == "" {
// 		t.Errorf("Battle ID has not been initialized")
// 	}

// 	if gotUpdateSessionNewBattleID != battle.ID {
// 		t.Errorf("got wrong new battle ID in UpdateSession: %q, want %q", gotUpdateSessionNewBattleID, battle.ID)
// 	}
// }

// func TestRedisBattlesService_ListBattles(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test")
// 	}
// 	mockSessionService := &sessions.SessionServiceMock{}
// 	mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
// 		return &sessions.Session{}, nil
// 	}
// 	mockSessionService.UpdateFn = func(s *sessions.Session) error {
// 		return nil
// 	}

// 	redisContainer, err := test.SetupRedisTestContainer(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer redisContainer.Terminate(ctx)
// 	client := redis.NewClient(&redis.Options{
// 		Addr: redisContainer.Addr,
// 	})

// 	r, err := redis_battle_service.New(client, mockSessionService)
// 	if err != nil {
// 		t.Fatalf("Could not get Redis Battle Service: %v", err)
// 	}

// 	battle := &battles.Battle{
// 		Name:  "somename",
// 		Owner: "someowner",
// 		Rank:  9,
// 		Status: battles.BattleStatus{
// 			BattleID:  "mybattleid",
// 			Status:    "hello world",
// 			Timestamp: "testtimestamp",
// 		},
// 		Members: []string{"foo", "bar"},
// 	}
// 	_ = r.CreateBattle(battle)

// 	gotBattleList, err := r.ListBattles()
// 	if err != nil {
// 		t.Fatalf("Could not get battle list: %v", err)
// 	}

// 	wantBattleList := []battles.Battle{*battle}

// 	if !reflect.DeepEqual(gotBattleList, wantBattleList) {
// 		t.Errorf("Got wrong battle list. Got:\n%+v\n\nWant:\n%+v\n", gotBattleList, wantBattleList)
// 	}
// }

// func TestBattleService_JoinBattle(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test")
// 	}
// 	redisContainer, err := test.SetupRedisTestContainer(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer redisContainer.Terminate(ctx)
// 	client := redis.NewClient(&redis.Options{
// 		Addr: redisContainer.Addr,
// 	})

// 	Convey("Given a battleservice", t, func() {
// 		mockSessionService := &sessions.SessionServiceMock{}
// 		r, err := redis_battle_service.New(client, mockSessionService)
// 		if err != nil {
// 			t.Fatalf("Could not get Redis Battle Service: %v", err)
// 		}

// 		Convey("Given a session that does not exist", func() {
// 			sessionID := "dummy-404-session"
// 			battleID := "dummy-battle"

// 			var gotSessionIDParam string
// 			mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
// 				gotSessionIDParam = id
// 				return nil, fmt.Errorf("could not find the session")
// 			}

// 			Convey("When JoinBattle is called", func() {
// 				_, err := r.JoinBattle(sessionID, battleID)
// 				So(gotSessionIDParam, ShouldEqual, sessionID)

// 				Convey("Then an error is returned", func() {
// 					So(err, ShouldBeError)
// 				})
// 			})
// 		})

// 		Convey("Given a session that exists", func() {
// 			sessionID := "a-session-id-that-exists"
// 			battleID := "some-battle"
// 			var gotNewBattleID string
// 			mockSessionService.GetSessionFn = func(id string) (*sessions.Session, error) {
// 				return &sessions.Session{ID: sessionID}, nil
// 			}
// 			mockSessionService.UpdateFn = func(s *sessions.Session) error {
// 				gotNewBattleID = s.BattleID
// 				return nil
// 			}
// 			testBattle := &battles.Battle{
// 				Owner: "some-battle-owner",
// 			}
// 			err := r.CreateBattle(testBattle)
// 			So(err, ShouldBeNil)

// 			Convey("When JoinBattle is called", func() {
// 				updatedBattleData, err := r.JoinBattle(sessionID, battleID)

// 				Convey("There is no error returned", func() {
// 					So(err, ShouldBeNil)
// 				})

// 				Convey("Then the session is added to the battle members", func() {
// 					So(updatedBattleData, ShouldNotBeNil)
// 					So(sessionID, ShouldBeIn, updatedBattleData.Members)
// 				})

// 				Convey("Then the session's battle is updated to the new battle", func() {
// 					So(gotNewBattleID, ShouldEqual, battleID)
// 				})
// 			})
// 		})
// 	})
// }
