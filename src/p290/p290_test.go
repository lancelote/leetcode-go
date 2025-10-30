package p290

import (
	"testing"
)

func Test_wordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		{
			"match",
			"abba",
			"dog cat cat dog",
			true,
		},
		{
			"no match",
			"abba",
			"dog cat cat fish",
			false,
		},
		{
			"no match - same byte",
			"aaaa",
			"dog cat cat dog",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordPattern(tt.pattern, tt.s)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
