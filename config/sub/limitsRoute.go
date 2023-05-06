package sub

type RouteRateLimit struct {
	Guild   RateLimitOptions `json:"guild"`
	Webhook RateLimitOptions `json:"webhook"`
	Channel RateLimitOptions `json:"channel"`
}
