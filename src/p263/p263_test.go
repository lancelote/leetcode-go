package p263

import (
	"testing"
)

func Test_isUgly(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			"negative",
			-1,
			false,
		},
		{
			"zero",
			0,
			false,
		},
		{
			"one",
			1,
			true,
		},
		{
			"ugly",
			6,
			true,
		},
		{
			"not ugly",
			14,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUgly(tt.n)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
