package sub

type ChannelLimits struct {
	MaxPins     int `json:"maxPins"`
	MaxTopic    int `json:"maxTopic"`
	MaxWebhooks int `json:"maxWebhooks"`
}
