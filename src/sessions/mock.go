package sessions

type SessionServiceMock struct {
	LoginFn      func(nick string) (sessionID string, err error)
	GetSessionFn func(id string) (*Session, error)
	UpdateFn     func(*Session) error
}

func (s *SessionServiceMock) Login(nick string) (sessionID string, err error) {
	return s.LoginFn(nick)
}

func (s *SessionServiceMock) GetSession(id string) (*Session, error) {
	return s.GetSessionFn(id)
}

func (s *SessionServiceMock) Update(session *Session) error {
	return s.UpdateFn(session)
}
