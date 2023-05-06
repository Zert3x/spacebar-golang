package config

import (
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type DefaultsConfiguration struct {
	Guild sub.GuildDefaults `json:"guild"`
	User  sub.UserDefaults  `json:"user"`
}

func init() {
	viper.SetDefault("defaults.guild.maxPresences", 250000)
	viper.SetDefault("defaults.guild.maxVideoChannelUsers", 200)
	viper.SetDefault("defaults.guild.afkTimeout", 300)
	viper.SetDefault("defaults.guild.defaultMessageNotifications", 1)
	viper.SetDefault("defaults.user.premium", true)
	viper.SetDefault("defaults.user.premiumType", 2)
	viper.SetDefault("defaults.user.verified", true)
}
