package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			input: []string{"abc", "cba", "xyz", "zyx"},
			expected: map[string][]string{
				"abc": {"abc", "cba"},
				"xyz": {"xyz", "zyx"},
			},
		},
	}

	for _, test := range tests {
		result := findAnagrams(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Input: %v\nExpected: %v\nGot: %v\n", test.input, test.expected, result)
		}
	}
}
