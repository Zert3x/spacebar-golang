package config

type Config struct {
	Gateway       EndpointConfiguration       `json:"gateway"`
	Cdn           CdnConfiguration            `json:"cdn"`
	Api           ApiConfiguration            `json:"api"`
	General       GeneralConfiguration        `json:"general"`
	Limits        LimitConfiguration          `json:"limits"`
	Security      SecurityConfiguration       `json:"security"`
	Login         LoginConfiguration          `json:"login"`
	Register      RegisterConfiguration       `json:"register"`
	Regions       RegionConfiguration         `json:"regions"`
	Guild         GuildConfiguration          `json:"guild"`
	Gif           GifConfiguration            `json:"gif"`
	RabbitMQ      RabbitMQConfiguration       `json:"rabbitmq"`
	Kafka         KafkaConfiguration          `json:"kafka"`
	Templates     TemplateConfiguration       `json:"templates"`
	Metrics       MetricsConfiguration        `json:"metrics"`
	Sentry        SentryConfiguration         `json:"sentry"`
	Defaults      DefaultsConfiguration       `json:"defaults"`
	External      ExternalTokensConfiguration `json:"external"`
	Email         EmailConfiguration          `json:"email"`
	PasswordReset PasswordResetConfiguration  `json:"passwordReset"`
}
