package web_frontend

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/tobyjwebb/teamchess/src/settings"
	user_service "github.com/tobyjwebb/teamchess/src/user/service"
	redis_user_service "github.com/tobyjwebb/teamchess/src/user/service/redis"
)

type Server struct {
	config      settings.Config
	UserService user_service.UserService
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

func (s *Server) Start() {
	s.router = chi.NewRouter()

	s.initUserService()
	s.setupHtmlHandler()
	s.setupRoutes()

	log.Println("Starting server on", s.config.FrontendAddr)
	http.ListenAndServe(s.config.FrontendAddr, s.router)
}

func (s *Server) initUserService() {
	if s.UserService != nil {
		return
	}
	client, err := s.getRedisClient()
	if err != nil {
		panic(err)
	}
	redisUserService, err := redis_user_service.New(client)
	if err != nil {
		panic(err)
	}
	s.UserService = redisUserService
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
