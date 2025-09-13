package p1143

import (
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			"partial match",
			"abcde",
			"ace",
			3,
		},
		{
			"full match",
			"abc",
			"abc",
			3,
		},
		{
			"no match",
			"abc",
			"def",
			0,
		},
		{
			"random1",
			"bsbininm",
			"jmjkbkjkv",
			1,
		},
		{
			"random2",
			"dknkdizqxkdczafixidorgfcnkrirmhmzqbcfuvojsxwraxe",
			"dulixqfgvipenkfubgtyxujixspoxmhgvahqdmzmlyhajerqz",
			14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
