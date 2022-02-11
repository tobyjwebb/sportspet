package web_frontend

import (
	"log"
	"net/http"
)

func (s *Server) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	log.Printf("In LoginHandler: r: %v\n", r)
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var url string
	nick := r.FormValue("nick")
	if nick == "" {
		url = "/nick-required.html"
	} else {
		sessionID, err := s.UserService.Login(nick)
		if err != nil {
			panic(err)
		}
		url = "/web/session/" + sessionID
	}

	http.Redirect(rw, r, url, http.StatusTemporaryRedirect)
}
