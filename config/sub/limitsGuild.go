package sub

type GuildLimits struct {
	MaxRoles              int `json:"maxRoles"`
	MaxEmojis             int `json:"maxEmojis"`
	MaxMembers            int `json:"maxMembers"`
	MaxChannels           int `json:"maxChannels"`
	MaxChannelsInCategory int `json:"maxChannelsInCategory"`
}
