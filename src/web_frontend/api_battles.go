package web_frontend

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupBattlesRoutes() *chi.Mux {
	challenges := chi.NewRouter()
	challenges.Get("/{challenge_id}/state", s.getBatleStateHandler)
	return challenges
}

func (s *Server) getBatleStateHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement getBatleStateHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{
		"board": "          (XXX 64-chars, one for each pos in board)             ",
		"turn":"white",
		"latest_logs": [
			{"n": 5, "who":"user1", "piece":"q", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 4, "who":"user2", "piece":"P", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 3, "who":"user6", "piece":"k", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"}
		]
	}`)
}
