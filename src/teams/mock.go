package teams

type TeamServiceMock struct {
	CreateTeamFn     func(name, sessionID string) (teamID string, err error)
	CreateTeamCalled bool
}

func (t *TeamServiceMock) CreateTeam(name, sessionID string) (teamID string, err error) {
	t.CreateTeamCalled = true
	if t.CreateTeamFn != nil {
		return t.CreateTeamFn(name, sessionID)
	}
	return "", nil
}
