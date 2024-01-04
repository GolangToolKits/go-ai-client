package goopenaiclient

import (
	"encoding/json"
	"log"
)

// Client for OpenAI
type Client struct {
	OrgKey  string
	APIKey  string
	Version string
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

// ChatRequestUnmarshal ChatRequestUnmarshal
func ChatRequestUnmarshal(jsonData string) *ChatRequest {
	var rtn ChatRequest
	var jsonBlob = []byte(jsonData)
	err := json.Unmarshal(jsonBlob, &rtn)
	if err != nil {
		log.Println("JSON unmarshal error: ", err)
	}
	return &rtn
}
