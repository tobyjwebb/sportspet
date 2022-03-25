package web_frontend

import (
	"net/http"
)

func (s *Server) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	// log.Printf("In LoginHandler: r: %v\n", r)
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var url string
	nick := r.FormValue("nick")
	if nick == "" {
		url = "/nick-required.html"
	} else {
		sessionID, err := s.SessionService.Login(nick)
		if err != nil {
			panic(err)
		}
		if sessionID == "" {
			url = "/nick-already-used.html" + sessionID
		} else {
			url = "/main.html?session=" + sessionID
		}
	}

	http.Redirect(rw, r, url, http.StatusTemporaryRedirect)
}
