package config

import "github.com/spf13/viper"

type CdnConfiguration struct {
	ResizeHeightMax int     `json:"resizeHeightMax"`
	ResizeWidthMax  int     `json:"resizeWidthMax"`
	ImagorServerUrl *string `json:"imagorServerUrl"`
	EndpointPublic  *string `json:"endpointPublic"`
	EndpointPrivate *string `json:"endpointPrivate"`
}

func init() {
	viper.SetDefault("cdn.resizeHeightMax", 1000)
	viper.SetDefault("cdn.resizeWidthMax", 1000)
}
