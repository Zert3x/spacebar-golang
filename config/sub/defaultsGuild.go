package sub

type GuildDefaults struct {
	MaxPresences                int `json:"maxPresences"`
	MaxVideoChannelUsers        int `json:"maxVideoChannelUsers"`
	AfkTimeout                  int `json:"afkTimeout"`
	DefaultMessageNotifications int `json:"defaultMessageNotifications"`
	ExplicitContentFilter       int `json:"explicitContentFilter"`
}

func init() {

}
