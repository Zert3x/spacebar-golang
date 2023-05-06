package sub

type PasswordConfiguration struct {
	Required     bool `json:"required"`
	MinLength    int  `json:"minLength"`
	MinNumbers   int  `json:"minNumbers"`
	MinUpperCase int  `json:"minUpperCase"`
	MinSymbols   int  `json:"minSymbols"`
}
