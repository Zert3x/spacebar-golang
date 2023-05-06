package sub

type GlobalRateLimit struct {
	Limit   int  `json:"limit"`
	Window  int  `json:"window"`
	Enabled bool `json:"enabled"`
}

type GlobalRateLimits struct {
	Register    GlobalRateLimit `json:"register"`
	SendMessage GlobalRateLimit `json:"sendMessage"`
}
