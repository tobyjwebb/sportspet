package service

import "net/http"

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`Hello there!`))
}
