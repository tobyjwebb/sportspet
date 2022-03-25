package sessions

type SessionServiceMock struct {
	LoginFn func(nick string) (sessionID string, err error)
}

func (s *SessionServiceMock) Login(nick string) (sessionID string, err error) {
	return s.LoginFn(nick)
}
