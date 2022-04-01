package redis_test

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tobyjwebb/teamchess/src/challenges"
	redis_challenge_service "github.com/tobyjwebb/teamchess/src/challenges/redis"
	"github.com/tobyjwebb/teamchess/src/test"
)

func TestChallengeService_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_challenge_service.New(client)
	if err != nil {
		t.Fatalf("Could not get Redis challenge Service: %v", err)
	}

	Convey("Given a challenge", t, func() {
		c := &challenges.Challenge{}

		Convey("When Create is called", func() {
			err := r.Create(c)

			Convey("The ID is filled in", func() {
				So(c.ID, ShouldNotBeEmpty)
			})
			Convey("There is no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestChallengeService_List(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx := context.Background()

	redisContainer, err := test.SetupRedisTestContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)
	client := redis.NewClient(&redis.Options{
		Addr: redisContainer.Addr,
	})

	r, err := redis_challenge_service.New(client)
	if err != nil {
		t.Fatalf("Could not get Redis challenge Service: %v", err)
	}

	Convey("Given a service", t, func() {
		Convey("When List is called", func() {
			list, err := r.List("")

			Convey("The result is empty", func() {
				So(list, ShouldBeEmpty)
			})
			Convey("There is no error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("Given a challenge", func() {
			challengerID := "testChallenger"
			challengeeID := "testChallengee"
			c := &challenges.Challenge{
				ChallengerTeamID: challengerID,
				ChallengeeTeamID: challengeeID,
			}
			err := r.Create(c)
			Convey("No error", func() {
				So(err, ShouldBeNil)
			})

			Convey("When List is called for the challenger team", func() {
				list, err := r.List(challengerID)

				Convey("The result is not empty, and should return the new challenge", func() {
					So(list, ShouldNotBeEmpty)
					challenge := findChallengeByID(list, c.ID)
					So(challenge, ShouldNotBeNil)
					Convey("The challenger fields are as expected", func() {
						So(challenge.ChallengerTeamID, ShouldEqual, challengerID)
						So(challenge.ChallengeeTeamID, ShouldEqual, challengeeID)
					})
				})
				Convey("There is no error", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When List is called for the challengee team", func() {
				list, err := r.List(challengeeID)

				Convey("The result is not empty, and should return the new challenge", func() {
					So(list, ShouldNotBeEmpty)
					challenge := findChallengeByID(list, c.ID)
					So(challenge, ShouldNotBeNil)
					Convey("The challenger fields are as expected", func() {
						So(challenge.ChallengerTeamID, ShouldEqual, challengerID)
						So(challenge.ChallengeeTeamID, ShouldEqual, challengeeID)
					})
				})
				Convey("There is no error", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When List is called for another team", func() {
				list, err := r.List("foo-team")

				Convey("The result is empty", func() {
					So(list, ShouldBeEmpty)
				})
				Convey("There is no error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func findChallengeByID(list []challenges.Challenge, id string) *challenges.Challenge {
	for _, c := range list {
		if c.ID == id {
			return &c
		}
	}
	return nil
}
