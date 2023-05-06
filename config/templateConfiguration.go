package config

import "github.com/spf13/viper"

type TemplateConfiguration struct {
	Enabled               bool `json:"enabled"`
	AllowTemplateCreation bool `json:"allowTemplateCreation"`
	AllowDiscordTemplates bool `json:"allowDiscordTemplates"`
	AllowRaws             bool `json:"allowRaws"`
}

func init() {
	viper.SetDefault("template.enabled", true)
	viper.SetDefault("template.allowTemplateCreation", true)
	viper.SetDefault("template.allowDiscordTemplates", true)
	viper.SetDefault("template.allowRaws", true)
}
