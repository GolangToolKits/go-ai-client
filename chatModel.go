package goopenaiclient

import "time"

// ChatRequest ChatRequest
type ChatRequest struct {
	Model            string         `json:"model"`
	Messages         []any          `json:"messages"`          //system, user, assistant, tool or function message
	FrequencyPenalty any            `json:"frequency_penalty"` //num or null
	LogitBias        any            `json:"logit_bias"`        //map or null
	Logprobs         any            `json:"logprobs"`          //boolean or null
	TopLogprobs      any            `json:"top_logprobs"`      //nteger or null
	MaxTokens        any            `json:"max_tokens"`        //integer or null
	NumberGenTokens  any            `json:"n"`                 //integer or null
	PresencePenalty  any            `json:"presence_penalty"`  //number or null
	ResponseFormat   ChatRespFormat `json:"response_format"`
	Seed             any            `json:"seed"`        //integer or null
	Stop             any            `json:"stop"`        //string / array / null
	Stream           any            `json:"stream"`      //boolean or null
	Temperature      any            `json:"temperature"` //number or null
	TopProbability   any            `json:"top_p"`       //number or null
	Tools            []ChatTool     `json:"tools"`
	ToolChoice       any            `json:"tool_choice"` //string or chatTool
	User             string
}

// ChatResponse ChatResponse
type ChatResponse struct {
	ID                string       `json:"id"`
	Object            string       `json:"object"`
	Created           time.Time    `json:"created"`
	Model             string       `json:"model"`
	Choices           []ChatChoice `json:"choices"`
	Usage             ChatUsage    `json:"usage"`
	SystemFingerprint string       `json:"system_fingerprint"`
}

// ChatTool ChatTool
type ChatTool struct {
	Type     string       `json:"type"`
	Function ChatFunction `json:"function"`
}

// ChatFunction ChatFunction
type ChatFunction struct {
	Name        string `json:"name"`
	Description any    `json:"description"` //string or null
	Parameters  any    `json:"parameters"`  //ChatParameters or null
}

// ChatParameters ChatParameters
type ChatParameters struct {
	Type       string         `json:"type"`
	Properties ChatProperties `json:"properties"`
	Required   []string       `json:"required"`
}

// ChatProperties ChatProperties
type ChatProperties struct {
	Location ChatLocation `json:"location"`
	Unit     ChatUnit     `json:"unit"`
}

// ChatUnit ChatUnit
type ChatUnit struct {
	Type string   `json:"type"`
	Enum []string `json:"enum"`
}

// ChatLocation Location
type ChatLocation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ChatRespFormat ChatRespFormat
type ChatRespFormat struct {
	Type string `json:"type"`
}

// ChatImageURL ChatImageURL
type ChatImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail"`
}

// ChatTextContentParts ChatContent
type ChatTextContentParts struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ChatImageContentParts ChatImageContent
type ChatImageContentParts struct {
	Type     string       `json:"type"`
	ImageURL ChatImageURL `json:"image_url"`
}

// ChatSystemMessage ChatSystemMessage
type ChatSystemMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

// ChatUserMessage ChatUserMessage
type ChatUserMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"` //string or array of parts
	Name    string `json:"name"`
}

// ChatAssistantMessage ChatAssistantMessage
type ChatAssistantMessage struct {
	Role      string       `json:"role"`
	Content   any          `json:"content"` //string or null
	Name      string       `json:"name"`
	ToolCalls ChatToolCall `json:"tool_calls"`
}

// ChatFunctionMessage ChatFunctionMessage
type ChatFunctionMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"` //string or null
	Name    string `json:"name"`
}

// ChatToolMessage ChatToolMessage
type ChatToolMessage struct {
	Role       string `json:"role"`
	Content    string `json:"content"`
	ToolCallID string `json:"tool_call_id"`
}

// ChatToolCall ChatToolCall
type ChatToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function ChatToolFunction `json:"function"`
}

// ChatToolFunction ChatToolFunction
type ChatToolFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ChatMessage Message
type ChatMessage struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
	Name    string `json:"name"`
}

// ChatLogprobContent LogprobContent
type ChatLogprobContent struct {
	Token       string               `json:"token"`
	Logprob     float64              `json:"logprob"`
	Bytes       []byte               `json:"bytes"`
	TopLogprobs []ChatLogprobContent `json:"top_logprobs"`
}

// ChatLogprob Logprob
type ChatLogprob struct {
	Content []ChatLogprobContent `json:"content"`
}

// ChatChoice Choice
type ChatChoice struct {
	Index        int64       `json:"index"`
	Message      ChatMessage `json:"message"`
	Logprobs     ChatLogprob `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

// ChatUsage ChatUsage
type ChatUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}
