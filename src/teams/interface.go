package teams

type TeamsService interface {
	CreateTeam(name, sessionID string) (teamID string, err error)
}
