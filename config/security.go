package config

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/spf13/viper"
	"spacebar-server/config/sub"
)

type SecurityConfiguration struct {
	Captcha                            sub.CaptchaConfiguration   `json:"captcha"`
	TwoFactor                          sub.TwoFactorConfiguration `json:"twoFactor"`
	AutoUpdate                         bool                       `json:"autoUpdate"`
	RequestSignature                   string                     `json:"requestSignature"`
	JwtSecret                          string                     `json:"jwtSecret"`
	ForwardedFor                       *string                    `json:"forwardedFor"`
	IpdataApiKey                       *string                    `json:"ipdataApiKey"`
	MfaBackupCodeCount                 int                        `json:"mfaBackupCodeCount"`
	StatsWorldReadable                 bool                       `json:"statsWorldReadable"`
	DefaultRegistrationTokenExpiration int                        `json:"defaultRegistrationTokenExpiration"`
}

func init() {
	viper.SetDefault("security.twoFactor.generateBackupCodes", true)
	sig := make([]byte, 32)
	rand.Read(sig)
	jwt := make([]byte, 256)
	rand.Read(jwt)
	viper.SetDefault("security.requestSignature", base64.StdEncoding.EncodeToString(sig))
	viper.SetDefault("security.jwtSecret", base64.StdEncoding.EncodeToString(jwt))
	viper.SetDefault("security.mfaBackupCodeCount", 10)
	viper.SetDefault("security.statsWorldReadable", true)
	viper.SetDefault("security.defaultRegistrationTokenExpiration", 1000*60*60*24*7)
}
