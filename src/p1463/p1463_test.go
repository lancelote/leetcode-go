package p1463

import (
	"testing"
)

func Test_longestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"whole array with one deletion",
			[]int{1, 1, 0, 1},
			3,
		},
		{
			"part of the array with one deletion",
			[]int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			5,
		},
		{
			"only ones",
			[]int{1, 1, 1},
			2,
		},
		{
			"no ones",
			[]int{0, 0, 0},
			0,
		},
		{
			"starts with one",
			[]int{1, 0, 0},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestSubarray(tt.nums)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
