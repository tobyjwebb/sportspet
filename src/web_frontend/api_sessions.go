package web_frontend

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupSessionsRoutes() *chi.Mux {
	sessions := chi.NewRouter()
	sessions.Get("/me", s.getSessionHandler)
	return sessions
}

func (s *Server) getSessionHandler(rw http.ResponseWriter, r *http.Request) {
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
	res := sessionResponse{}
	if session.TeamID != "" {
		team, err := s.TeamService.GetTeamData(session.TeamID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Battle = team.Status.BattleID
		res.Team = sessionTeamResponse{
			ID:   team.ID,
			Name: team.Name,
		}
	}
	setJSON(rw)
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type sessionResponse struct {
	Battle string              `json:"battle"`
	Team   sessionTeamResponse `json:"team"`
}
type sessionTeamResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
