package goopenaiclient

// OpenAI OpenAI client
type OpenAI interface {
	Completions(r *ChatRequest) *ChatResponse
}

// go mod init github.com/GolangToolKits/go-openai-client
