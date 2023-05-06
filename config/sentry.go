package config

import (
	"github.com/spf13/viper"
	"os"
)

type SentryConfiguration struct {
	Enabled         bool    `json:"enabled"`
	Endpoint        string  `json:"endpoint"`
	TraceSampleRate float32 `json:"traceSampleRate"`
	Environment     string  `json:"environment"`
}

func init() {
	viper.SetDefault("sentry.traceSampleRate", 1.0)
	hostname, _ := os.Hostname()
	viper.SetDefault("sentry.environment", hostname)
}
