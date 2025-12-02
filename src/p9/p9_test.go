package p9

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			"basic true",
			121,
			true,
		},
		{
			"negative",
			-121,
			false,
		},
		{
			"basic false",
			10,
			false,
		},
		{
			"long with zeros",
			1000021,
			false,
		},
		{
			"possible zero division",
			1001,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
