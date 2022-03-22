package web_frontend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/tobyjwebb/teamchess/src/teams"
)

func setJSON(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func (s *Server) setupTeamsRoutes() *chi.Mux {
	teams := chi.NewRouter()
	teams.Get("/", s.ListTeamsHandler)
	teams.Post("/", s.CreateTeamHandler)
	teams.Post("/{team_id}/join", s.JoinTeamHandler)
	teams.Post("/{team_id}/leave", s.leaveTeam)
	return teams
}

func (s *Server) ListTeamsHandler(rw http.ResponseWriter, r *http.Request) {
	teams, err := s.TeamService.ListTeams()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	setJSON(rw)
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(teams); err != nil {
		panic(err)
	}
}

func (s *Server) CreateTeamHandler(rw http.ResponseWriter, r *http.Request) {
	owner := r.FormValue("owner")
	team := &teams.Team{
		Name:    r.FormValue("name"),
		Owner:   owner,
		Members: []string{owner},
	}

	if err := s.TeamService.CreateTeam(team); err != nil {
		log.Printf("Error creating team: %v", err)
		panic(err)
	}

	setJSON(rw)
	rw.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(team); err != nil {
		panic(err)
	}
}

func getSessionIDFromAuth(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 || authSplit[0] != "Bearer" {
		return ""
	}
	return authSplit[1]
}

func (s *Server) JoinTeamHandler(rw http.ResponseWriter, r *http.Request) {
	sessionID := getSessionIDFromAuth(r)
	teamIDParam := chi.URLParam(r, "team_id")
	log.Printf("team id: %q", teamIDParam)
	if _, err := s.TeamService.JoinTeam(sessionID, teamIDParam); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("error joining team: %v", err)
		return
	}
}

func (s *Server) leaveTeam(rw http.ResponseWriter, r *http.Request) {
	// XXX implement leaveTeam action
	setJSON(rw)
	fmt.Fprintf(rw, `{"warning":"not implemented"}`)
}
