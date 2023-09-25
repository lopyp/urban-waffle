package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"\\n", "", true},
		{"\\2n", "2n", false},
		{"\\43m", "444m", false},
		{"\\\\n", "\\n", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := unpackString(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("Expected error: %v, but got error: %v", test.wantErr, err)
			}

			if result != test.expected {
				t.Errorf("Expected: %s, but got: %s", test.expected, result)
			}
		})
	}
}

