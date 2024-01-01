package goopenaiclient

import "time"

// ChatImageURL ChatImageURL
type ChatImageURL struct {
	URL string `json:"url"`
}

// ChatContent ChatContent
type ChatContent struct {
	Type     string       `json:"type"`
	Text     string       `json:"text"`
	ImageURL ChatImageURL `json:"image_url"`
}

// Message Message
type Message struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
}

// ChatRequest ChatRequest
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// LogprobContent LogprobContent
type LogprobContent struct {
	Token       string           `json:"token"`
	Logprob     float64          `json:"logprob"`
	Bytes       []byte           `json:"bytes"`
	TopLogprobs []LogprobContent `json:"top_logprobs"`
}

// Logprob Logprob
type Logprob struct {
	Content []LogprobContent `json:"content"`
}

// Choice Choice
type Choice struct {
	Index        int64   `json:"index"`
	Message      Message `json:"message"`
	Logprobs     Logprob `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

// ChatUsage ChatUsage
type ChatUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// ChatResponse ChatResponse
type ChatResponse struct {
	ID                string    `json:"id"`
	Object            string    `json:"object"`
	Created           time.Time `json:"created"`
	Model             string    `json:"model"`
	Choices           []Choice  `json:"choices"`
	Usage             ChatUsage `json:"usage"`
	SystemFingerprint string    `json:"system_fingerprint"`
}

// Completions Completions
func (c *Client) Completions(r *ChatRequest) *ChatResponse {
	var rtn ChatResponse

	return &rtn
}
