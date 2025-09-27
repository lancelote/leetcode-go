package p88

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			"first",
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			"second",
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
		{
			"third",
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if diff := cmp.Diff(tt.want, tt.nums1); diff != "" {
				t.Error(diff)
			}
		})
	}
}
