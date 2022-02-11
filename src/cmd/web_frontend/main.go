package main

import (
	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func main() {
	cfg := settings.GetConfig()
	server := web_frontend.NewServer(cfg)
	server.Start()
}
