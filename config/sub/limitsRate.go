package sub

type RateLimits struct {
	Enabled bool             `json:"enabled"`
	Global  RateLimitOptions `json:"global"`
	Error   RateLimitOptions `json:"error"`
	Routes  RouteRateLimit   `json:"routes"`
}
