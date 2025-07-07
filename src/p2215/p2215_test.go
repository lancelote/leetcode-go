package p2215

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func toSet(s []int) map[int]struct{} {
	set := make(map[int]struct{})

	for _, x := range s {
		set[x] = struct{}{}
	}

	return set
}

func Test_findDifference(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{
			"both have unique",
			[]int{1, 2, 3},
			[]int{2, 4, 6},
			[][]int{{1, 3}, {4, 6}},
		},
		{
			"only first has unique",
			[]int{1, 2, 3, 3},
			[]int{1, 1, 2, 2},
			[][]int{{3}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findDifference(tt.nums1, tt.nums2)

			only1 := actual[0]
			only2 := actual[1]

			want1 := tt.want[0]
			want2 := tt.want[1]

			if diff := cmp.Diff(toSet(want1), toSet(only1)); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(toSet(want2), toSet(only2)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
