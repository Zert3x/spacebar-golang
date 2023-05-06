package sub

type AuthRateLimit struct {
	Login    RateLimitOptions `json:"login"`
	Register RateLimitOptions `json:"register"`
}
