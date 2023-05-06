package config

import (
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type RegisterConfiguration struct {
	Email                      sub.RegistrationEmailConfiguration `json:"email"`
	DateOfBirth                sub.DateOfBirthConfiguration       `json:"dateOfBirth"`
	Password                   sub.PasswordConfiguration          `json:"password"`
	Disabled                   bool                               `json:"disabled"`
	RequireCaptcha             bool                               `json:"requireCaptcha"`
	RequireInvite              bool                               `json:"requireInvite"`
	GuestsRequireInvite        bool                               `json:"guestsRequireInvite"`
	AllowNewRegistration       bool                               `json:"allowNewRegistration"`
	AllowMultipleAccounts      bool                               `json:"allowMultipleAccounts"`
	BlockProxies               bool                               `json:"blockProxies"`
	IncrementingDiscriminators bool                               `json:"incrementingDiscriminators"`
	DefaultRights              string                             `json:"defaultRights"`
}

func init() {
	viper.SetDefault("register.email.blocklist", true)
	viper.SetDefault("register.dateOfBirth.required", true)
	viper.SetDefault("register.dateOfBirth.minimum", 13)
	viper.SetDefault("register.password.minLength", 8)
	viper.SetDefault("register.password.minNumbers", 2)
	viper.SetDefault("register.password.minSymbols", 2)
	viper.SetDefault("register.requireCaptcha", true)
	viper.SetDefault("register.guestsRequireInvite", true)
	viper.SetDefault("register.allowNewRegistration", true)
	viper.SetDefault("register.allowMultipleAccounts", true)
	viper.SetDefault("register.blockProxies", true)
	viper.SetDefault("register.defaultRights", "875069521787904")
}
