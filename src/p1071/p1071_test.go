package p1071

import (
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			"whole second",
			"ABCABC",
			"ABC",
			"ABC",
		},
		{
			"substring of both",
			"ABABAB",
			"ABAB",
			"AB",
		},
		{
			"nothing",
			"LEET",
			"CODE",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcdOfStrings(tt.str1, tt.str2)
			if result != tt.want {
				t.Errorf("gcdOfStrings(%s, %s) = %s, want %s", tt.str1, tt.str2, result, tt.want)
			}
		})
	}
}
