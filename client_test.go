package goopenaiclient

import (
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
