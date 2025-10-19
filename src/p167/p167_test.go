package p167

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			"first",
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			"second",
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			"third",
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
		{
			"fourth",
			[]int{5, 25, 75},
			100,
			[]int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
