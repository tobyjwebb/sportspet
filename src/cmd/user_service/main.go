package main

import (
	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/user/service"
)

func main() {
	cfg := settings.GetConfig()
	server := service.NewServer(cfg)
	server.Start()
}
