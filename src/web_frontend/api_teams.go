package web_frontend

import (
	"fmt"
	"net/http"
)

func (s *Server) TeamsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	switch r.Method {
	case http.MethodPost:
		fmt.Fprintf(rw, `{"name":"new_team_stub","id":"stub-team-id"}`)
	case http.MethodGet:
		fmt.Fprintf(rw, `[
			{"name":"team1"},
			{"name":"team2"},
			{"name":"team3"}
			]`)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(rw, `{"error": "method not allowed"}`)
		return
	}
}
