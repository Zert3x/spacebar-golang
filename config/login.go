package config

import "github.com/spf13/viper"

type LoginConfiguration struct {
	RequireCaptcha      bool `json:"requireCaptcha"`
	RequireVerification bool `json:"requireVerification"`
}

func init() {
	viper.SetDefault("login.requireCaptcha", false)
	viper.SetDefault("login.requireVerification", false)
}
