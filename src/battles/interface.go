package battles

type BattleService interface {
	Create(battle *Battle) error
	GetData(id string) (*Battle, error)
	Update(battle *Battle) error
}
