package sub

type RateLimitOptions struct {
	Bot    *int  `json:"bot"`
	Count  int   `json:"count"`
	Window int   `json:"windows"`
	OnlyIp *bool `json:"onlyIp"`
}
