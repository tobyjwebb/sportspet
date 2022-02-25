package web_frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/teams"
	user_service "github.com/tobyjwebb/teamchess/src/user/service"
	redis_user_service "github.com/tobyjwebb/teamchess/src/user/service/redis"
)

type Server struct {
	config      settings.Config
	UserService user_service.UserService
	TeamService teams.TeamService
	redisClient *redis.Client
	router      *chi.Mux
}

func NewServer(c *settings.Config) *Server {
	config := c
	if config == nil {
		config = settings.GetConfig()
	}
	return &Server{
		config: *config,
	}
}

func (s *Server) Start() error {
	s.router = chi.NewRouter()

	if err := s.initUserService(); err != nil {
		return fmt.Errorf("could not init user service: %w", err)
	}
	s.setupHtmlHandler()
	s.setupRoutes()

	log.Println("Starting server on", s.config.FrontendAddr)
	return http.ListenAndServe(s.config.FrontendAddr, s.router)
}

func (s *Server) initUserService() error {
	if s.UserService != nil {
		return nil
	}
	client, err := s.getRedisClient()
	if err != nil {
		return fmt.Errorf("could not init Redis client: %w", err)
	}
	redisUserService, err := redis_user_service.New(client)
	if err != nil {
		return fmt.Errorf("could not init Redis user service: %w", err)
	}
	s.UserService = redisUserService
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
