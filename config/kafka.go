package config

import "spacebar-server/config/sub"

type KafkaConfiguration struct {
	Brokers []sub.KafkaBroker `json:"brokers"`
}
