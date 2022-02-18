package web_frontend

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func setJSON(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func (s *Server) setupTeamsRoutes() *chi.Mux {
	teams := chi.NewRouter()
	teams.Get("/", s.listTeams)
	teams.Post("/", s.createTeam)
	teams.Post("/{team_id}/join", s.joinTeam)
	return teams
}

func (s *Server) listTeams(rw http.ResponseWriter, r *http.Request) {
	setJSON(rw)
	fmt.Fprintf(rw, `[
			{"name":"team1","id":"id1"},
			{"name":"team2","id":"id2"},
			{"name":"team2.5","id":"id2andahalf"},
			{"name":"team3","id":"id3"}
			]`)
}

func (s *Server) createTeam(rw http.ResponseWriter, r *http.Request) {
	setJSON(rw)
	fmt.Fprintf(rw, `{"name":"new_team_stub","id":"stub-team-id"}`)
}

func (s *Server) joinTeam(rw http.ResponseWriter, r *http.Request) {
	setJSON(rw)
	fmt.Fprintf(rw, `{"warning":"not implemented"}`)
}
