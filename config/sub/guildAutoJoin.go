package sub

type AutoJoinConfiguration struct {
	Enabled  bool     `json:"enabled"`
	Guilds   []string `json:"guilds"`
	CanLeave bool     `json:"canLeave"`
}
