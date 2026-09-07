package client

import (
	"testing"
)

func TestReplaceSensitiveHeadersObjectScale(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Redact Authorization header",
			input:    []byte("Authorization: Basic YWRtaW46cGFzc3dvcmQ="),
			expected: "Authorization: ******",
		},
		{
			name:     "Redact X-SDS-AUTH-TOKEN header",
			input:    []byte("X-SDS-AUTH-TOKEN: abc123xyz789"),
			expected: "X-SDS-AUTH-TOKEN: ******",
		},
		{
			name:     "Redact Cookie auth_cookie",
			input:    []byte("Cookie: auth_cookie=secret123; Path=/"),
			expected: "Cookie: auth_cookie=******; Path=/",
		},
		{
			name:     "Multiple sensitive headers",
			input:    []byte("Authorization: Basic YWRtaW46cGFzc3dvcmQ=\nX-SDS-AUTH-TOKEN: abc123"),
			expected: "Authorization: ******\nX-SDS-AUTH-TOKEN: ******",
		},
		{
			name:     "No sensitive headers",
			input:    []byte("Content-Type: application/json"),
			expected: "Content-Type: application/json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replaceSensitiveHeadersObjectScale(tt.input)
			if result != tt.expected {
				t.Errorf("replaceSensitiveHeadersObjectScale() = %v, want %v", result, tt.expected)
			}
		})
	}
}