package p1456

import (
	"testing"
)

func Test_maxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			"mix",
			"abciiidef",
			3,
			3,
		},
		{
			"only vowels",
			"aeiou",
			2,
			2,
		},
		{
			"not enough",
			"leetcode",
			3,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxVowels(tt.s, tt.k)
			if actual != tt.want {
				t.Errorf("got %d, want %d", actual, tt.want)
			}
		})
	}
}
