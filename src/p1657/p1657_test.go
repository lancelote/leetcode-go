package p1657

import (
	"testing"
)

func Test_closeStrings(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		{
			"just swap",
			"abc",
			"bca",
			true,
		},
		{
			"impossible",
			"a",
			"aa",
			false,
		},
		{
			"swap and change",
			"cabbba",
			"abbccc",
			true,
		},
		{
			"different runes",
			"uau",
			"ssx",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closeStrings(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
