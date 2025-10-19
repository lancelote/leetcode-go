package p58

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"first",
			"Hello World",
			5,
		},
		{
			"second",
			"   fly me   to   the moon  ",
			4,
		},
		{
			"third",
			"luffy is still joyboy",
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLastWord(tt.s)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
