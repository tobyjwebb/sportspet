package battles

type BattleServiceMock struct {
	CreateBattleFn  func(*Battle) error
	GetBattleDataFn func(id string) (*Battle, error)
	UpdateFn        func(battle *Battle) error
}

func (t *BattleServiceMock) Create(battle *Battle) error {
	return t.CreateBattleFn(battle)
}
func (t *BattleServiceMock) GetData(id string) (*Battle, error) {
	return t.GetBattleDataFn(id)
}
func (t *BattleServiceMock) Update(battle *Battle) error {
	return t.UpdateFn(battle)
}
