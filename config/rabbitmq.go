package config

import "github.com/spf13/viper"

type RabbitMQConfiguration struct {
	Host *string `json:"host"`
}

func init() {
	viper.SetDefault("rabbitmq.host", nil)
}
