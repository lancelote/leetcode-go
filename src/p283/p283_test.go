package p283

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// func moveZeroes(nums []int)  {
func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"has something to move",
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			"has nothing to move",
			[]int{0},
			[]int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if diff := cmp.Diff(tt.want, tt.nums); diff != "" {
				t.Error(diff)
			}
		})
	}
}
