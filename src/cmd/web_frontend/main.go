package main

import (
	"fmt"

	"github.com/tobyjwebb/teamchess/src/settings"
	"github.com/tobyjwebb/teamchess/src/web_frontend"
)

func main() {
	cfg := settings.GetConfig()
	server := web_frontend.NewServer(cfg)
	if err := server.Start(); err != nil {
		fmt.Println("Could not start server:", err)
	}
}
