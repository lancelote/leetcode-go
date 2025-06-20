package p1768

import (
	"testing"
)

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			"same length",
			"abc",
			"pqr",
			"apbqcr",
		},
		{
			"first is shorter",
			"ab",
			"pqrs",
			"apbqrs",
		},
		{
			"second is shorter",
			"abcd",
			"pq",
			"apbqcd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeAlternately(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("mergeAlternately(%s, %s) = %s, want %s", tt.word1, tt.word2, result, tt.expected)
			}
		})
	}
}
