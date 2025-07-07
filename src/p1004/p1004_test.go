package p1004

import (
	"testing"
)

func Test_longestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			"last",
			[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			2,
			6,
		},
		{
			"first",
			[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			3,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestOnes(tt.nums, tt.k)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
