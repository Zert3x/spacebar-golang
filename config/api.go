package config

import "github.com/spf13/viper"

type ApiConfiguration struct {
	DefaultVersion string   `json:"defaultVersion"`
	ActiveVersions []string `json:"activeVersions"`
	EndpointPublic *string  `json:"endpointPublic"`
}

func init() {
	viper.SetDefault("api.defaultVersion", "9")
	viper.SetDefault("api.activeVersions", []string{"6", "7", "8", "9"})
	viper.SetDefault("api.endpointPublic", nil)
}
