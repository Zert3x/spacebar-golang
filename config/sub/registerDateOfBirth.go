package sub

type DateOfBirthConfiguration struct {
	Required bool `json:"required"`
	Minimum  int  `json:"minimum"`
}
