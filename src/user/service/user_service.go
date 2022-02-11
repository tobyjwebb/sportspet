package user_service

type UserService interface {
	Login(nick string) (sessionID string, err error)
}
