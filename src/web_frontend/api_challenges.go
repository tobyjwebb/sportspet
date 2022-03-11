package web_frontend

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupChallengesRoutes() *chi.Mux {
	challenges := chi.NewRouter()
	challenges.Get("/", s.getSessionChallengesHandler)
	challenges.Post("/", s.CreateChallengeHandler)
	challenges.Post("/{challenge_id}/accept", s.AcceptChallengeHandler)
	return challenges
}

func (s *Server) getSessionChallengesHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement getSessionChallengesHandler
	// XXX Requires Bearer token
	setJSON(rw)
	fmt.Fprintf(rw, `[
		{
			"id": "aaabbb-cccc-ffff-11122233",
			"challenger": {
				"id": "aaabbb-cccc-ffff-11122233",
				"name": "The Fooers"
			}
			"timestamp": "2000-12-05T12:34:56Z"
		}
	]`)
}

func (s *Server) CreateChallengeHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement CreateChallengeHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{"id":"aaabbb-cccc-ffff-11122233"}`)
}

func (s *Server) AcceptChallengeHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement AcceptChallengeHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{"battle_id":"aaabbb-cccc-ffff-11122233"}`)
}
