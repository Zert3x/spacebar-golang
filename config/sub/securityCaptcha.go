package sub

type CaptchaConfiguration struct {
	Enabled bool    `json:"enabled"`
	Service *string `json:"service"`
	SiteKey *string `json:"siteKey"`
	Secret  *string `json:"secret"`
}
