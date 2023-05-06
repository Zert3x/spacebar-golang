package sub

type Region struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Location struct {
		Latitude  int `json:"latitude"`
		Longitude int `json:"longitude"`
	} `json:"location"`
	Vip        bool `json:"vip"`
	Custom     bool `json:"custom"`
	Deprecated bool `json:"deprecated"`
}
