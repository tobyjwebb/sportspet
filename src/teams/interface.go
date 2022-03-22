package teams

type TeamService interface {
	CreateTeam(*Team) error
	ListTeams() ([]Team, error)
	// JoinTeam(sessionID, teamID string) (*Team, error)
}
