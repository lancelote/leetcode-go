package p80

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			"first",
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		{
			"second",
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			"not enough",
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if tt.k != got {
				t.Fatalf("want %d, got %d", tt.k, got)
			}
			if diff := cmp.Diff(tt.want, tt.nums[:got]); diff != "" {
				t.Error(diff)
			}
		})
	}
}
