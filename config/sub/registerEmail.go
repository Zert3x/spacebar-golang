package sub

type RegistrationEmailConfiguration struct {
	Required  bool     `json:"required"`
	Allowlist bool     `json:"allowlist"`
	Blocklist bool     `json:"blocklist"`
	Domains   []string `json:"domains"`
}
