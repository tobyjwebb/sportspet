package teams

import "fmt"

type TeamServiceMock struct {
	CreateTeamFn  func(*Team) error
	ListTeamsFn   func() ([]Team, error)
	JoinTeamFn    func(sessionID, teamID string) (*Team, error)
	GetTeamDataFn func(id string) (*Team, error)
	UpdateFn      func(*Team) error
}

func (t *TeamServiceMock) CreateTeam(team *Team) error {
	if t.CreateTeamFn != nil {
		return t.CreateTeamFn(team)
	}
	return fmt.Errorf("CreateTeamFn has not been defined")
}

func (t *TeamServiceMock) ListTeams() ([]Team, error) {
	if t.ListTeamsFn != nil {
		return t.ListTeamsFn()
	}
	return nil, fmt.Errorf("ListTeamsFn has not been defined")
}

func (t *TeamServiceMock) JoinTeam(sessionID, teamID string) (*Team, error) {
	if t.JoinTeamFn != nil {
		return t.JoinTeamFn(sessionID, teamID)
	}
	return nil, fmt.Errorf("JoinTeamFn has not been defined")
}

func (t *TeamServiceMock) GetTeamData(id string) (*Team, error) {
	return t.GetTeamDataFn(id)
}

func (t *TeamServiceMock) Update(team *Team) error {
	return t.UpdateFn(team)
}
