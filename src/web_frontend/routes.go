package web_frontend

import "github.com/go-chi/chi/v5"

func (s *Server) setupRoutes() {
	s.router.Post("/login", s.LoginHandler)
	s.router.Mount("/api/v1", s.setupAPIRoutes())
}

func (s *Server) setupAPIRoutes() *chi.Mux {
	api := chi.NewRouter()
	api.Mount("/teams", s.setupTeamsRoutes())
	return api
}
