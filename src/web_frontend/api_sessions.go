package web_frontend

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupSessionsRoutes() *chi.Mux {
	sessions := chi.NewRouter()
	sessions.Get("/{challenge_id}", s.getSessionHandler)
	return sessions
}

func (s *Server) getSessionHandler(rw http.ResponseWriter, r *http.Request) {
	// XXX implement getSessionHandler
	setJSON(rw)
	fmt.Fprintf(rw, `{
		"battle": "aabbcc-dd-11-33322323232233",
		"team": {
			"id": "9999999ff-12332-23k4234j233",
			"name": "Qux"
		}
	}`)
}
