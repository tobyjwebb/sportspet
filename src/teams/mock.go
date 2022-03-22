package teams

import "fmt"

type TeamServiceMock struct {
	CreateTeamFn func(*Team) error
	ListTeamsFn  func() ([]Team, error)
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
