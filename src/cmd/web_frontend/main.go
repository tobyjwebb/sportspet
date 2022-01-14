package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("TC_FRONTEND_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	fmt.Println("Starting server on", addr)
	http.ListenAndServe(addr, nil)
}
