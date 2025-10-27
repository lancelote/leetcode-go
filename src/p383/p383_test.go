package p383

import (
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			"single letter false",
			"a",
			"b",
			false,
		},
		{
			"one missing letter",
			"aa",
			"ab",
			false,
		},
		{
			"enough letters",
			"aa",
			"aab",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
