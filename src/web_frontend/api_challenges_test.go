package web_frontend_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tobyjwebb/teamchess/src/challenges"
	"github.com/tobyjwebb/teamchess/src/sessions"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func TestChallengeTeamHandler(t *testing.T) {
	Convey("Given a server with mocked services", t, func() {
		server := web_frontend.NewServer(nil)
		challengServiceMock := &challenges.ChallengeServiceMock{}
		sessionServiceMock := &sessions.SessionServiceMock{}
		server.ChallengeService = challengServiceMock
		server.SessionService = sessionServiceMock
		response := httptest.NewRecorder()

		Convey("Given a request to challenge another team", func() {
			challengesURL := "/api/v1/challenges/"
			request, err := http.NewRequest(http.MethodPost, challengesURL, nil)
			So(err, ShouldBeNil)

			Convey("When the challenge request is sent with a missing authorization header", func() {
				server.ServeHTTP(response, request)

				Convey("Then the response is unauthorized", func() {
					So(response.Result().StatusCode, ShouldEqual, http.StatusUnauthorized)
				})
			})

			Convey("Given a valid authorization header", func() {
				userSessionID := "the-user-session-id"
				request.Header.Set("Authorization", "Bearer "+userSessionID)

				Convey("Given a missing teamid parameter", func() {
					Convey("When the request is sent", func() {
						server.ServeHTTP(response, request)
						Convey("Then the response is bad request", func() {
							So(response.Result().StatusCode, ShouldEqual, http.StatusBadRequest)
						})
					})
				})

				Convey("Given a teamid parameter", func() {
					testChallengedTeamID := "the-challenged-team-id"
					request.Form = url.Values{}
					request.Form.Add("team", testChallengedTeamID)
					Convey("Given a session service mock that returns an error", func() {
						sessionServiceMock.GetSessionFn = func(id string) (*sessions.Session, error) {
							return nil, fmt.Errorf("OhNo!")
						}
						Convey("When the request is sent", func() {
							server.ServeHTTP(response, request)
							Convey("Then the result is error", func() {
								So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
							})
						})
					})

					Convey("Given a session service mock that returns the session without a team ID", func() {
						var gotSessionID string
						sessionServiceMock.GetSessionFn = func(id string) (*sessions.Session, error) {
							gotSessionID = id
							return &sessions.Session{
								ID: id,
							}, nil
						}

						Convey("When the request is sent", func() {
							server.ServeHTTP(response, request)
							So(gotSessionID, ShouldEqual, userSessionID)
							Convey("Then the result is bad request", func() {
								So(response.Result().StatusCode, ShouldEqual, http.StatusBadRequest)
							})
						})
					})

					Convey("Given a session service mock that returns the session with a team ID", func() {
						testTeamID := "the-team-id"
						sessionServiceMock.GetSessionFn = func(id string) (*sessions.Session, error) {
							return &sessions.Session{
								ID:     id,
								TeamID: testTeamID,
							}, nil
						}

						Convey("Given a challenge mock which returns an error", func() {
							challengServiceMock.CreateFn = func(challenge *challenges.Challenge) error {
								return fmt.Errorf("testOHNO-Something-Bad-Happenned-Error")
							}

							Convey("When the challenge request is sent", func() {
								server.ServeHTTP(response, request)

								Convey("Then the result is error", func() {
									So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
								})

							})

						})

						Convey("Given a mock that returns success", func() {
							theChallengeID := "the-expected-challenge-id"
							var gotChallengerTeamID string
							var gotChallengedTeamID string
							challengServiceMock.CreateFn = func(challenge *challenges.Challenge) error {
								gotChallengerTeamID = challenge.ChallengerTeamID
								gotChallengedTeamID = challenge.ChallengeeTeamID
								challenge.ID = theChallengeID
								return nil
							}

							Convey("When the challenge request is sent", func() {
								server.ServeHTTP(response, request)

								So(gotChallengerTeamID, ShouldEqual, testTeamID)
								So(gotChallengedTeamID, ShouldEqual, testChallengedTeamID)

								Convey("Then created challenge-id is returned", func() {
									So(response.Result().StatusCode, ShouldEqual, http.StatusOK)
									gotJSON := &struct {
										ID string `json:"id"`
									}{}
									err := json.NewDecoder(response.Result().Body).Decode(gotJSON)
									So(err, ShouldBeNil)
									So(gotJSON.ID, ShouldEqual, theChallengeID)
								})
							})
						})
					})
				})

			})
		})
	})
}
