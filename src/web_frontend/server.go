package web_frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/challenges"
	"github.com/tobyjwebb/teamchess/src/sessions"
	session_service "github.com/tobyjwebb/teamchess/src/sessions"
	redis_session_service "github.com/tobyjwebb/teamchess/src/sessions/redis"
	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/teams"
	redis_team_service "github.com/tobyjwebb/teamchess/src/teams/redis"
)

type Server struct {
	config           settings.Config
	SessionService   session_service.SessionService
	TeamService      teams.TeamService
	ChallengeService challenges.ChallengeService
	redisClient      *redis.Client
	router           *chi.Mux
}

func NewServer(c *settings.Config) *Server {
	config := c
	if config == nil {
		config = settings.GetConfig()
	}
	s := &Server{
		config: *config,
		router: chi.NewRouter(),
	}
	s.mountHandlers()
	return s
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(rw, r)
}

func (s *Server) Start() error {
	if err := s.initSessionService(); err != nil {
		return fmt.Errorf("could not init user service: %w", err)
	}
	if err := s.initTeamService(s.SessionService); err != nil {
		return fmt.Errorf("could not init team service: %w", err)
	}

	log.Println("Starting server on", s.config.FrontendAddr)
	return http.ListenAndServe(s.config.FrontendAddr, s.router)
}

func (s *Server) initSessionService() error {
	if s.SessionService != nil {
		return nil
	}
	client, err := s.getRedisClient()
	if err != nil {
		return fmt.Errorf("could not init Redis client: %w", err)
	}
	redisSessionService, err := redis_session_service.New(client)
	if err != nil {
		return fmt.Errorf("could not init Redis user service: %w", err)
	}
	s.SessionService = redisSessionService
	return nil
}

func (s *Server) initTeamService(sessionService sessions.SessionService) error {
	if s.TeamService != nil {
		return nil
	}
	client, err := s.getRedisClient()
	if err != nil {
		return fmt.Errorf("could not init Redis client: %w", err)
	}
	redisTeamService, err := redis_team_service.New(client, sessionService)
	if err != nil {
		return fmt.Errorf("could not init Redis team service: %w", err)
	}
	s.TeamService = redisTeamService
	return nil
}

func (s *Server) getRedisClient() (*redis.Client, error) {
	if s.redisClient != nil {
		return s.redisClient, nil
	}
	client := redis.NewClient(&redis.Options{
		Addr: s.config.RedisAddr,
	})
	ctx := context.Background()
	res := client.Ping(ctx)
	if err := res.Err(); err != nil {
		return nil, err
	}
	s.redisClient = client
	return client, nil
}
