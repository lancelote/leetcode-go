package p189

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			"first",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"second",
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			"single element",
			[]int{1},
			2,
			[]int{1},
		},
		{
			"no elements",
			[]int{},
			42,
			[]int{},
		},
		{
			"big shift",
			[]int{1, 2},
			7,
			[]int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if diff := cmp.Diff(tt.want, tt.nums); diff != "" {
				t.Error(diff)
			}
		})
	}
}
