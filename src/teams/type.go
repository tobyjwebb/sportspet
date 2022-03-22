package teams

type Team struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Owner   string     `json:"owner,omitempty"` // stored as sessionID?
	Rank    int        `json:"rank"`
	Status  TeamStatus `json:"status"`
	Members []string   `json:"members,omitempty"`
}

type TeamStatus struct {
	BattleID  string `json:"battle_id,omitempty"`
	Status    string `json:"status,omitempty"`
	Timestamp string `json:"timestamp"`
}
