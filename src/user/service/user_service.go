package service

type UserService interface {
	Login(nick string) (sessionID string, err error)
}

// type server struct {
// 	config settings.Config
// }

// func NewServer(c *settings.Config) *server {
// 	return &server{config: *c}
// }

// func (s *server) Start() {
// 	http.HandleFunc("/api/v1/user/login", LoginHandler)

// 	addr := s.config.UserServiceAddr
// 	log.Println("Starting user service server on", addr)
// 	http.ListenAndServe(addr, nil)
// }
