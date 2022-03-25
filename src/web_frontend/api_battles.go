package web_frontend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const initialBoardStatus = "CHBQKBHCPPPPPPPP              qQ                ppppppppchbqkbhc"

func (s *Server) setupBattlesRoutes() *chi.Mux {
	challenges := chi.NewRouter()
	challenges.Get("/{challenge_id}/state", s.getBatleStateHandler)
	challenges.Get("/{challenge_id}/log", s.getBatleLogHandler)
	challenges.Post("/{challenge_id}/move", s.postBatleMoveHandler)
	return challenges
}

func (s *Server) getBatleStateHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement getBatleStateHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{
		"board": "%s",
		"white_team": "aabbcc-dd-11-33322323232233",
		"black_team": "9999999ff-12332-23k4234j233",
		"turn":"black",
		"latest_movements": [
			{"n": 5, "who":"user1", "piece":"q", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 4, "who":"user2", "piece":"P", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 3, "who":"user6", "piece":"k", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"}
		]
	}`, initialBoardStatus)
}

func (s *Server) getBatleLogHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement getBatleLogHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{"latest_movements":[
			{"n": 5, "who":"user1", "piece":"q", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 4, "who":"user2", "piece":"P", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
			{"n": 3, "who":"user6", "piece":"k", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"}
		]}`)
}

func (s *Server) postBatleMoveHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement postBatleMoveHandler
	from := r.FormValue("from")
	to := r.FormValue("to")
	log.Println("XXX move from", from, "to", to)
	setJSON(rw)
	fmt.Fprintf(rw, `{}`)
}
