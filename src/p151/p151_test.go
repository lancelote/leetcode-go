package p151

import (
	"testing"
)

// func reverseWords(s string) string {
func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"basic example",
			"the sky is blue",
			"blue is sky the",
		},
		{
			"spaces around",
			"  hello world  ",
			"world hello",
		},
		{
			"more space in between",
			"a good   example",
			"example good a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := reverseWords(tt.s)
			if actual != tt.want {
				t.Errorf("reverseWords(%v) = %v, want %v", tt.s, actual, tt.want)
			}
		})
	}
}
