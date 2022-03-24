package challenges

type Challenge struct {
	ID               string `json:"id"`
	ChallengerTeamID string `json:"challenging_team_id"`
	ChallengeeTeamID string `json:"challengee_team_id"`
}
