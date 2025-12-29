package p139

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
