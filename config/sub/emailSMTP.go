package sub

type SMTPConfiguration struct {
	Host     *string `json:"host"`
	Port     *int    `json:"port"`
	Secure   *bool   `json:"secure"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
