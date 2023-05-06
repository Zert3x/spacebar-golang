package database

import (
	"database/sql"
	"math/big"
	"time"
)

type User struct {
	Id                string         `json:"id" sql:"id"`
	Username          string         `json:"username" sql:"username"`
	Discriminator     string         `json:"discriminator" sql:"discriminator"`
	Avatar            sql.NullString `json:"avatar" sql:"avatar"`
	AccentColor       sql.NullString `json:"accent_color" sql:"accent_color"`
	Banner            sql.NullString `json:"banner" sql:"banner"`
	ThemeColors       sql.NullString `json:"theme_colors" sql:"theme_colors"`
	Pronouns          sql.NullString `json:"pronouns" sql:"pronouns"`
	Phone             sql.NullString `json:"phone" sql:"phone"`
	Desktop           bool           `json:"desktop" sql:"desktop"`
	Mobile            bool           `json:"mobile" sql:"mobile"`
	Premium           bool           `json:"premium" sql:"premium"`
	PremiumType       int            `json:"premium_type" sql:"premium_type"`
	Bot               bool           `json:"bot" sql:"bot"`
	Bio               string         `json:"bio" sql:"bio"`
	System            bool           `json:"system" sql:"system"`
	NsfwAllowed       bool           `json:"nsfw_allowed" sql:"nsfw_allowed"`
	MfaEnabled        bool           `json:"mfa_enabled" sql:"mfa_enabled"`
	WebauthnEnabled   bool           `json:"webauthn_enabled" sql:"webauthn_enabled"`
	TotpSecret        sql.NullString `json:"-" sql:"totp_secret"`
	TotpLastTicket    sql.NullString `json:"-" sql:"totp_last_ticket"`
	CreatedAt         time.Time      `json:"created_at" sql:"created_at"`
	PremiumSince      sql.NullTime   `json:"premium_since" sql:"premium_since"`
	Verified          bool           `json:"verified" sql:"verified"`
	Disabled          bool           `json:"disabled" sql:"disabled"`
	Deleted           bool           `json:"deleted" sql:"deleted"`
	Email             sql.NullString `json:"email" sql:"email"`
	Flags             string         `json:"flags" sql:"flags"` // TODO: Generate
	PublicFlags       int            `json:"public_flags" sql:"public_flags"`
	PurchasedFlags    int            `json:"purchased_flags" sql:"purchased_flags"`
	PremiumUsageFlags int            `json:"premium_usage_flags" sql:"premium_usage_flags"`
	Rights            *big.Int       `json:"rights" sql:"rights"`
	//Sessions          []Session      `json:"sessions" sql:"-"`
	RelationIds []string `json:"relationship_ids" sql:"relationship_ids"`
}

func (u *User) ToPublicUser() PublicUserProjection {
	var premiumSince *time.Time
	if u.PremiumSince.Valid {
		premiumSince = &u.PremiumSince.Time
	}

	return PublicUserProjection{
		Id:            u.Id,
		Username:      u.Username,
		Discriminator: u.Discriminator,
		PublicFlags:   u.PublicFlags,
		Avatar:        u.Avatar.String,
		AccentColor:   u.AccentColor.String,
		Banner:        u.Banner.String,
		Bio:           u.Bio,
		Bot:           u.Bot,
		PremiumSince:  premiumSince,
		PremiumType:   u.PremiumType,
		ThemeColors:   u.ThemeColors.String,
		Pronouns:      u.Pronouns.String,
	}
}

type PublicUserProjection struct {
	Id            string     `json:"id"`
	Username      string     `json:"username"`
	Discriminator string     `json:"discriminator"`
	PublicFlags   int        `json:"public_flags"`
	Avatar        string     `json:"avatar,omitempty"`
	AccentColor   string     `json:"accent_color"`
	Banner        string     `json:"banner,omitempty"`
	Bio           string     `json:"bio,omitempty"`
	Bot           bool       `json:"bot"`
	PremiumSince  *time.Time `json:"premium_since,omitempty"`
	PremiumType   int        `json:"premium_type"`
	ThemeColors   string     `json:"theme_colors"`
	Pronouns      string     `json:"pronouns"`
}
