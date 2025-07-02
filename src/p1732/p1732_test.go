package p1732

import (
	"testing"
)

// func largestAltitude(gain []int) int {
func Test_largestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{
			"first",
			[]int{-5, 1, 5, 0, -7},
			1,
		},
		{
			"second",
			[]int{-4, -3, -2, -1, 4, 3, 2},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := largestAltitude(tt.gain)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
