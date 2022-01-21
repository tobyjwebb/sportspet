package web_frontend

import (
	"log"
	"net/http"

	"github.com/tobyjwebb/teamchess/src/settings"
)

type server struct {
	config settings.Config
}

func NewServer(c *settings.Config) *server {
	return &server{config: *c}
}

func (s *server) Start() {
	setupHtmlHandler()

	log.Println("Starting server on", s.config.FrontendAddr)
	http.ListenAndServe(s.config.FrontendAddr, nil)
}
