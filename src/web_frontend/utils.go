package web_frontend

import (
	"net/http"
	"strings"
)

func getSessionIDFromAuth(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 || authSplit[0] != "Bearer" {
		return ""
	}
	return authSplit[1]
}
