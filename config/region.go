package config

import (
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type RegionConfiguration struct {
	Default             string       `json:"default"`
	UseDefaultAsOptimal bool         `json:"useDefaultAsOptimal"`
	Available           []sub.Region `json:"available"`
}

func init() {
	viper.SetDefault("region.default", "spacebar")
	viper.SetDefault("region.useDefaultAsOptimal", true)
	viper.SetDefault("region.available", []sub.Region{
		{
			Id:         "spacebar",
			Name:       "spacebar",
			Endpoint:   "127.0.0.1:3004",
			Vip:        false,
			Custom:     false,
			Deprecated: false,
		},
	})
}
