package config

import (
	"github.com/spf13/viper"
	"spacebar-server/types"
)

type GeneralConfiguration struct {
	InstanceName         string  `json:"instanceName"`
	InstanceDescription  string  `json:"instanceDescription"`
	FrontPage            *string `json:"frontPage"`
	TosPage              *string `json:"tosPage"`
	CorrespondenceEmail  *string `json:"correspondenceEmail"`
	CorrespondenceUserId *string `json:"correspondenceUserId"`
	Image                *string `json:"image"`
	InstanceId           string  `json:"instanceId"`
}

func init() {
	viper.SetDefault("general.instanceName", "Spacebar Instance")
	viper.SetDefault("general.instanceDescription", "This is a Spacebar instance made in the pre-release days")
	viper.SetDefault("general.instanceId", types.GenerateSnowflake(0, 0))
}
