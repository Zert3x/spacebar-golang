package config

import "github.com/spf13/viper"

type MetricsConfiguration struct {
	Timeout int `json:"timeout"`
}

func init() {
	viper.SetDefault("metrics.timeout", 30000)
}
