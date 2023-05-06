package config

import "github.com/spf13/viper"

type GifConfiguration struct {
	Enabled  bool    `json:"enabled"`
	Provider string  `json:"provider"`
	ApiKey   *string `json:"apiKey"`
}

func init() {
	viper.SetDefault("gif.enabled", true)
	viper.SetDefault("gif.provider", "tenor")
	viper.SetDefault("gif.apiKey", nil)
}
