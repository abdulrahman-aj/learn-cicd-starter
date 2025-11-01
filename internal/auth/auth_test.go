package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers        http.Header
		expectedAPIKey string
		expectedErr    error
	}{
		"valid API key": {
			headers:        http.Header{"Authorization": []string{"ApiKey xyz"}},
			expectedAPIKey: "xyz",
		},
		"invalid API key": {
			headers:     http.Header{},
			expectedErr: auth.ErrNoAuthHeaderIncluded,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			APIKey, err := auth.GetAPIKey(tt.headers)
			if APIKey != tt.expectedAPIKey {
				t.Fatalf("expected: %v, got: %v", tt.expectedAPIKey, APIKey)
			}
			if err != tt.expectedErr {
				t.Fatalf("expected err: %v, got: %v", tt.expectedErr, err)
			}
		})
	}
}
