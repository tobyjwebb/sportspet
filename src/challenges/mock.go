package challenges

type ChallengeServiceMock struct {
	CreateFn func(challenge *Challenge) error
	ListFn   func(teamID string) ([]Challenge, error)
	DeleteFn func(challengeID string) error
	ReadFn   func(id string) (*Challenge, error)
}

func (t *ChallengeServiceMock) Create(challenge *Challenge) error {
	return t.CreateFn(challenge)
}

func (t *ChallengeServiceMock) List(teamID string) ([]Challenge, error) {
	return t.ListFn(teamID)
}

func (t *ChallengeServiceMock) Delete(challengeID string) error {
	return t.DeleteFn(challengeID)
}

func (t *ChallengeServiceMock) Read(id string) (*Challenge, error) {
	return t.ReadFn(id)
}
