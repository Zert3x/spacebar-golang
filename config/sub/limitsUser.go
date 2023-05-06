package sub

type UserLimits struct {
	MaxGuilds   int `json:"maxGuilds"`
	MaxUsername int `json:"maxUsername"`
	MaxFriends  int `json:"maxFriends"`
}
