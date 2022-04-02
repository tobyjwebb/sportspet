package web_frontend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tobyjwebb/teamchess/src/battles"
	"github.com/tobyjwebb/teamchess/src/board"
)

func (s *Server) setupBattlesRoutes() *chi.Mux {
	challenges := chi.NewRouter()
	challenges.Get("/{battle_id}/state", s.getBatleStateHandler)
	challenges.Get("/{battle_id}/log", s.getBatleLogHandler)
	challenges.Post("/{battle_id}/move", s.postBatleMoveHandler)
	return challenges
}

func (s *Server) getBatleStateHandler(rw http.ResponseWriter, r *http.Request) {
	battleID := chi.URLParam(r, "battle_id")
	log.Println("Got battle id", battleID)

	status, battle, err := s.doGetBattleState(battleID)
	if err != nil {
		rw.WriteHeader(status)
		return
	}

	// TODO: Implement latest movements?:
	// "latest_movements": [
	// 	{"n": 5, "who":"user1", "piece":"q", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
	// 	{"n": 4, "who":"user2", "piece":"P", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"},
	// 	{"n": 3, "who":"user6", "piece":"k", "from": "A5", "to":"C6", "timestamp":"2022-02-22T11:11:11Z"}
	// ]

	turn := "black"
	if battle.MoveCount%2 == 0 {
		turn = "white"
	}
	setJSON(rw)
	fmt.Fprintf(rw, `{
		"board": "%s",
		"white_team": "%s",
		"black_team": "%s",
		"turn":"%s"
	}`, battle.Board, battle.WhiteTeamID, battle.BlackTeamID, turn)
}

func (s *Server) doGetBattleState(battleID string) (status int, battle *battles.Battle, err error) {
	b, err := s.BattleService.GetData(battleID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, b, nil
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
	log.Println("Move from", from, "to", to)
	battleID := chi.URLParam(r, "battle_id")
	log.Println("Got battle id", battleID)
	sessionID := getSessionIDFromAuth(r)
	status, err := s.doMovePiece(sessionID, battleID, from, to)
	rw.WriteHeader(status)
	if err != nil {
		return
	}
	setJSON(rw)
	fmt.Fprintf(rw, `{}`)
}

func (s *Server) doMovePiece(sessionID, battleID, from, to string) (status int, err error) {
	// TODO: Check if allowed to move (teams's turn)

	battle, err := s.BattleService.GetData(battleID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	b := board.New(battle.Board)
	b.MovePiece(from, to)
	after := b.State()

	battle.Board = after
	battle.MoveCount++
	if err := s.BattleService.Update(battle); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}
