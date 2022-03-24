package challenges

type ChallengeService interface {
	Create(challenge *Challenge) error
	List(teamID string) ([]Challenge, error)
}
