package config

type EndpointConfiguration struct {
	EndpointClient  *string `json:"endpointClient"`
	EndpointPrivate *string `json:"endpointPrivate"`
	EndpointPublic  *string `json:"endpointPublic"`
}
