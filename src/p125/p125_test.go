package p125

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			"first",
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"second",
			"race a car",
			false,
		},
		{
			"third",
			" ",
			true,
		},
		{
			"forth",
			"0P",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
