package config

import "github.com/spf13/viper"

type PasswordResetConfiguration struct {
	RequireCaptcha bool `json:"requireCaptcha"`
}

func init() {
	viper.SetDefault("passwordReset.requireCaptcha", false)
}
