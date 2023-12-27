package goopenaiclient

// Client for OpenAI
type Client struct {
	OrgKey string
	APIKey string
}

// NewClient NewClient
func NewClient(key string, keyType string) OpenAI {
	var c Client
	if keyType == OrgKey {
		c.OrgKey = key
	} else if keyType == APIKey {
		c.APIKey = key
	}
	return &c
}
