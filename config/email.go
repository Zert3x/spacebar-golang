package config

import "spacebar-server/config/sub"

type EmailConfiguration struct {
	Provider *string                   `json:"provider"`
	Smtp     sub.SMTPConfiguration     `json:"smtp"`
	Mailgun  sub.MailGunConfiguration  `json:"mailgun"`
	Mailjet  sub.MailJetConfiguration  `json:"mailjet"`
	Sendgrid sub.SendGridConfiguration `json:"sendgrid"`
}
