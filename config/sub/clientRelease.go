package sub

type ClientReleaseConfiguration struct {
	UseLocalRelease bool   `json:"useLocalRelease"`
	UpstreamVersion string `json:"upstreamVersion"`
}
