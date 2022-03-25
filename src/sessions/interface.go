package sessions

type SessionService interface {
	// Login logs a user in, returning a sessionID, or empty string if the nick is already in use
	Login(nick string) (sessionID string, err error)
}
