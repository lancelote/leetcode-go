package p875

import (
	"testing"
)

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{
			"first",
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			"second",
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			"third",
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
