package sub

type MessageLimits struct {
	MaxCharacters        int `json:"maxCharacters"`
	MaxTTSCharacters     int `json:"maxTTSCharacters"`
	MaxReactions         int `json:"maxReactions"`
	MaxAttachmentSize    int `json:"maxAttachmentSize"`
	MaxBulkDelete        int `json:"maxBulkDelete"`
	MaxEmbedDownloadSize int `json:"maxEmbedDownloadSize"`
}
