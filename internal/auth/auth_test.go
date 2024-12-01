package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		wantErr     bool
	}{
		{
			name: "Correct Headers",
			headers: http.Header{
				"Authorization": {"ApiKey testKey"},
			},
			expectedKey: "testKey",
			wantErr:     true,
		},
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			wantErr:     true,
		},
		{
			name: "Authorization Header does not have enough parts",
			headers: http.Header{
				"Authorization": {"ApiKey"},
			},
			expectedKey: "",
			wantErr:     true,
		},
		{
			name: "Authorization Header does not start with ApiKey",
			headers: http.Header{
				"Authorization": {"Bearer test"},
			},
			expectedKey: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, expected %v", key, tt.expectedKey)
			}
		})
	}
}
