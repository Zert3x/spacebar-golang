package config

import (
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type LimitConfiguration struct {
	User         sub.UserLimits       `json:"user"`
	Guild        sub.GuildLimits      `json:"guild"`
	Message      sub.MessageLimits    `json:"message"`
	Channel      sub.ChannelLimits    `json:"channel"`
	Rate         sub.RateLimits       `json:"rate"`
	AbsoluteRate sub.GlobalRateLimits `json:"absoluteRate"`
}

func init() {
	viper.SetDefault("limits.user.maxGuilds", 1048576)
	viper.SetDefault("limits.user.maxUsername", 32)
	viper.SetDefault("limits.user.maxFriends", 5000)

	viper.SetDefault("limits.guild.maxRoles", 1000)
	viper.SetDefault("limits.guild.maxEmojis", 2000)
	viper.SetDefault("limits.guild.maxMembers", 25000000)
	viper.SetDefault("limits.guild.maxChannels", 65535)
	viper.SetDefault("limits.guild.maxChannelsInCategory", 65535)

	viper.SetDefault("limits.message.maxCharacters", 1048576)
	viper.SetDefault("limits.message.maxTTSCharacters", 160)
	viper.SetDefault("limits.message.maxReactions", 2048)
	viper.SetDefault("limits.message.maxAttachmentSize", 1024*1024*1024)
	viper.SetDefault("limits.message.maxBulkDelete", 1000)
	viper.SetDefault("limits.message.maxEmbedDownloadSize", 1024*1024*5)

	viper.SetDefault("limits.channel.maxPins", 500)
	viper.SetDefault("limits.channel.maxTopic", 1024)
	viper.SetDefault("limits.channel.maxWebhooks", 100)

	viper.SetDefault("limits.rate.enabled", false)
	viper.SetDefault("limits.rate.ip.count", 500)
	viper.SetDefault("limits.rate.ip.window", 5)
	viper.SetDefault("limits.rate.global.count", 250)
	viper.SetDefault("limits.rate.global.window", 5)
	viper.SetDefault("limits.rate.error.count", 10)
	viper.SetDefault("limits.rate.error.window", 5)

	viper.SetDefault("limits.rates.routes.guild.count", 5)
	viper.SetDefault("limits.rates.routes.guild.window", 5)
	viper.SetDefault("limits.rates.routes.webhook.count", 5)
	viper.SetDefault("limits.rates.routes.webhook.window", 5)
	viper.SetDefault("limits.rates.routes.channel.count", 5)
	viper.SetDefault("limits.rates.routes.channel.window", 5)
	viper.SetDefault("limits.rates.routes.auth.login.count", 5)
	viper.SetDefault("limits.rates.routes.auth.login.window", 60)
	viper.SetDefault("limits.rates.routes.auth.register.count", 2)
	viper.SetDefault("limits.rates.routes.auth.register.count", 60*60*12)

	viper.SetDefault("limits.absoluteRate.register.limit", 25)
	viper.SetDefault("limits.absoluteRate.register.window", 60*60*1000)
	viper.SetDefault("limits.absoluteRate.register.enabled", true)
	viper.SetDefault("limits.absoluteRate.sendMessage.limit", 200)
	viper.SetDefault("limits.absoluteRate.sendMessage.window", 60*1000)
	viper.SetDefault("limits.absoluteRate.sendMessage.enabled", true)
}
