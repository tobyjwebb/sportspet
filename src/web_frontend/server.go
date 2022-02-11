package web_frontend

import (
	"log"
	"net/http"

	"github.com/tobyjwebb/teamchess/src/settings"
	user_service "github.com/tobyjwebb/teamchess/src/user/service"
	"github.com/tobyjwebb/teamchess/src/user/service/redis"
)

type Server struct {
	config      settings.Config
	UserService user_service.UserService
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
	s.initUserService()
	setupHtmlHandler()
	s.SetupRoutes()

	log.Println("Starting server on", s.config.FrontendAddr)
	http.ListenAndServe(s.config.FrontendAddr, nil)
}

func (s *Server) SetupRoutes() {
	http.HandleFunc("/login", s.LoginHandler)
}

func (s *Server) initUserService() {
	if s.UserService != nil {
		return
	}
	redisUserService, err := redis.New(s.config.RedisAddr)
	if err != nil {
		panic(err)
	}
	s.UserService = redisUserService
}
