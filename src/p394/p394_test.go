package p394

import (
	"testing"
)

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"basic",
			"3[a]2[bc]",
			"aaabcbc",
		},
		{
			"nested",
			"3[a2[c]]",
			"accaccacc",
		},
		{
			"string with no number",
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeString(tt.s)
			if got != tt.want {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
