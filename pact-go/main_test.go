package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email  string
		result bool
	}{
		{"test@example.com", true},
		{"test.email+regex@example.com", true},
		{"test.email@example.co.uk", true},
		{"test@exa_mple.com", false},
		{"test", false},
		{"test@.com", false},
		{"test@com", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			if got != tt.result {
				t.Errorf("IsValidEmail(%s) = %v; want %v", tt.email, got, tt.result)
			}
		})
	}
}
