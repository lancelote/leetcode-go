package p724

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"in the middle",
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			"doesn't exist",
			[]int{1, 2, 3},
			-1,
		},
		{
			"zero",
			[]int{2, 1, -1},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pivotIndex(tt.nums)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
