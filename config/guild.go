package config

import (
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type GuildConfiguration struct {
	Discovery       sub.DiscoveryConfiguration `json:"discovery"`
	AutoJoin        sub.AutoJoinConfiguration  `json:"autoJoin"`
	DefaultFeatures []string                   `json:"defaultFeatures"`
}

func init() {
	viper.SetDefault("guild.discovery.limit", 24)
	viper.SetDefault("guild.autoJoin.enabled", true)
	viper.SetDefault("guild.autoJoin.canLeave", true)
}
