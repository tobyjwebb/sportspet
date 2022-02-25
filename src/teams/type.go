package teams

type Team struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Owner   string   `json:"owner,omitempty"` // stored as sessionID?
	Members []string `json:"members,omitempty"`
}
