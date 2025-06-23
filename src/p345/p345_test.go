package p345

import (
	"testing"
)

// func reverseVowels(s string) string {
func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"uppercase",
			"IceCreAm",
			"AceCreIm",
		},
		{
			"lowercase",
			"leetcode",
			"leotcede",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := reverseVowels(tt.s)
			if actual != tt.want {
				t.Errorf("want %v, got %v", tt.want, actual)
			}
		})
	}
}
