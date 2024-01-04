package goopenaiclient

import (
	//"encoding/json"
	//"fmt"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		key     string
		keyType string
	}
	tests := []struct {
		name string
		args args
		want OpenAI
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				key:     "1234",
				keyType: "org",
			},
			want: &Client{
				OrgKey: "1234",
			},
		},
		{
			name: "test 2",
			args: args{
				key:     "12345",
				keyType: "api",
			},
			want: &Client{
				APIKey: "12345",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.key, tt.args.keyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatRequestUnmarshal(t *testing.T) {
	var cm ChatMessage

	var tcm ChatTextContentParts
	tcm.Type = "text"
	tcm.Text = "What’s in this image?"

	var iurl ChatImageURL
	iurl.URL = "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"
	var icm ChatImageContentParts
	icm.Type = "image_url"
	icm.ImageURL = iurl

	var cont = []any{tcm, icm}
	cm.Content = cont

	type args struct {
		jsonData string
	}
	tests := []struct {
		name string
		args args
		want *ChatRequest
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				jsonData: "{\r\n\"model\":\"gpt-3.5-turbo\",\r\n \"messages\":[\r\n {\r\n \"role\":\"system\",\r\n \"content\":\"You are a helpful assistant.\"\r\n},\r\n      {\r\n         \"role\":\"user\",\r\n         \"content\":\"Hello!\"\r\n      }\r\n   ]\r\n}				",
			},
			want: &ChatRequest{
				Model: "gpt-3.5-turbo",
				Messages: []ChatMessage{
					{
						Role:    "system",
						Content: "You are a helpful assistant.",
					},
					{
						Role:    "user",
						Content: "Hello!",
					},
				},
			},
		},
		{
			name: "test 2 image imput",
			args: args{
				jsonData: "{\r\n\"model\":\"gpt-4-vision-preview\",\r\n\"messages\":[\r\n{\r\n\"role\":\"user\",\r\n\"content\":[\r\n{\r\n\"type\": \"text\",\r\n\"text\": \"What\u2019s in this image?\"\r\n},\r\n{\r\n\"type\": \"image_url\",\r\n\"image_url\": {\r\n\"url\": \"https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg\"\r\n}\r\n}\r\n]\r\n}\r\n],\r\n\"max_tokens\": 300\r\n  }",
			},
			want: &ChatRequest{
				Model: "gpt-4-vision-preview",
				Messages: []ChatMessage{
					{
						Role: "user",
						Content: []any{
							map[string]string{
								"text": "What’s in this image?",
								"type": "text",
							},
							map[string]any{
								"image_url": map[string]any{
									"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg",
								},
								"type": "image_url",
							},
						},
					},
				},
				MaxTokens: 300,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChatRequestUnmarshal(tt.args.jsonData); !reflect.DeepEqual(got, tt.want) {
				//resJSON, _ := json.Marshal(got)
				//fmt.Println("json: ", string(resJSON))
				//fmt.Println("content: ", got.Messages[0].Content)
				t.Errorf("ChatRequestUnmarshal() = %v want %v", got, tt.want)
			}
		})
	}
}
