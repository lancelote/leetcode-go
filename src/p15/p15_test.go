package p15

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			"basic",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"no answer",
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			"single answer",
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			"duplicates",
			[]int{5, 5, 0, -5},
			[][]int{{-5, 0, 5}},
		},
		{
			"duplicate right",
			[]int{-5, 0, 5, 5},
			[][]int{{-5, 0, 5}},
		},
		{
			"duplicates left & right",
			[]int{-5, 0, 0, 5, 5},
			[][]int{{-5, 0, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
