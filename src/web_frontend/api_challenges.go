package web_frontend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tobyjwebb/teamchess/src/challenges"
)

func (s *Server) setupChallengesRoutes() *chi.Mux {
	challenges := chi.NewRouter()
	challenges.Get("/", s.getSessionChallengesHandler)
	challenges.Post("/", s.CreateChallengeHandler)
	challenges.Post("/{challenge_id}/accept", s.AcceptChallengeHandler)
	return challenges
}

func (s *Server) getSessionChallengesHandler(rw http.ResponseWriter, r *http.Request) {
	sessionID := getSessionIDFromAuth(r)
	if sessionID == "" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	session, err := s.SessionService.GetSession(sessionID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := []challengeResponse{}
	if session.TeamID != "" {
		challenges, err := s.ChallengeService.List(session.TeamID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, c := range challenges {
			// Check that we are not the challengers:
			if c.ChallengerTeamID != session.TeamID {
				res = append(res, challengeResponse{
					ID:        c.ID,
					Timestamp: "2000-12-05T12:34:56Z", // TODO implement
					Challenger: challengeChallengerResponse{
						ID: c.ChallengerTeamID,
					},
				})
			}
		}
	}

	// Fill in the team names:
	for _, c := range res {
		team, err := s.TeamService.GetTeamData(c.Challenger.ID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Challenger.Name = team.Name
	}
	setJSON(rw)
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type challengeResponse struct {
	ID         string                      `json:"id"`
	Challenger challengeChallengerResponse `json:"challenger"`
	Timestamp  string
}

type challengeChallengerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Server) CreateChallengeHandler(rw http.ResponseWriter, r *http.Request) {
	sessionID := getSessionIDFromAuth(r)
	if sessionID == "" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	challengedTeamID := r.FormValue("team")
	if challengedTeamID == "" {
		log.Println("Missing [team] url parameter")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := s.SessionService.GetSession(sessionID)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if session.TeamID == "" {
		log.Println("Session", sessionID, "has not joined a team yet")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check we haven't already challenged this team:
	if existing, err := s.ChallengeService.List(session.TeamID); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		for _, c := range existing {
			if c.ChallengeeTeamID == challengedTeamID || c.ChallengerTeamID == challengedTeamID {
				log.Println("Team", session.TeamID, "already challenged team", challengedTeamID)
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}

	challenge := &challenges.Challenge{
		ChallengerTeamID: session.TeamID,
		ChallengeeTeamID: challengedTeamID,
	}
	if err := s.ChallengeService.Create(challenge); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	setJSON(rw)
	if err := json.NewEncoder(rw).Encode(&createChallengeResponse{
		ID: challenge.ID,
	}); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) AcceptChallengeHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement AcceptChallengeHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{"battle_id":"aaabbb-cccc-ffff-11122233"}`)
}

type createChallengeResponse struct {
	ID string `json:"id"`
}
