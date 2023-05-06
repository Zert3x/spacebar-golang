package sub

type DiscoveryConfiguration struct {
	ShowAllGuilds     bool `json:"showAllGuilds"`
	UseRecommendation bool `json:"useRecommendation"`
	Offset            int  `json:"offset"`
	Limit             int  `json:"limit"`
}
