package web_frontend

import (
	"log"
	"net/http"

	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/user/service"
)

type Server struct {
	config      settings.Config
	UserService service.UserService
}

func NewServer(c *settings.Config) *Server {
	config := c
	if config == nil {
		config = settings.GetConfig()
	}
	return &Server{
		config: *config,
		UserService: &service.UserServiceMock{ // XXX use real user service
			LoginFn: func(nick string) (sessionID string, err error) {
				return "to-be-implemented-session-id-for-" + nick, nil
			},
		},
	}
}

func (s *Server) Start() {
	setupHtmlHandler()
	s.SetupRoutes()

	log.Println("Starting server on", s.config.FrontendAddr)
	http.ListenAndServe(s.config.FrontendAddr, nil)
}

func (s *Server) SetupRoutes() {
	http.HandleFunc("/login", s.LoginHandler)
}
