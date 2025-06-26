package p238

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// func productExceptSelf(nums []int) []int {
func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"basic example",
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			"with zero",
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := productExceptSelf(tt.nums)
			if diff := cmp.Diff(tt.want, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
