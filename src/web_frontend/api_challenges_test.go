package web_frontend_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tobyjwebb/teamchess/src/challenges"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func TestChallengeTeamHandler(t *testing.T) {
	Convey("Given a server with a ChallengeServiceMock", t, func() {
		server := web_frontend.NewServer(nil)
		mockService := &challenges.ChallengeServiceMock{}
		server.ChallengeService = mockService
		response := httptest.NewRecorder()

		Convey("Given a request to challenge another team", func() {
			url := "/api/v1/challenges/"
			request, err := http.NewRequest(http.MethodPost, url, nil)
			So(err, ShouldBeNil)

			Convey("When the challenge request is sent with a missing authorization header", func() {
				server.ServeHTTP(response, request)

				Convey("Then the response is unauthorized", func() {
					So(response.Result().StatusCode, ShouldEqual, http.StatusUnauthorized)
				})
			})

			Convey("Given a valid authorization header", func() {

				request.Header.Set("Authorization", "Bearer dummy-user-session-id")

				Convey("Given a mock which returns an error", func() {
					mockService.CreateFn = func(challenge *challenges.Challenge) error {
						return fmt.Errorf("testOHNO-Something-Bad-Happenned-Error")
					}

					Convey("When the challenge request is sent", func() {
						server.ServeHTTP(response, request)

						Convey("Then the result is error", func() {
							So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
						})

					})

				})

				// Convey("Given a mock that returns success", func() {

				// 	Convey("When the challenge request is sent", func() {

				// 		Convey("Then created challenge-id is returned", nil)

				// 	})

				// })

			})

		})

	})
}
