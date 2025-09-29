package p26

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		k    int
	}{
		{
			"single duplicate",
			[]int{1, 1, 2},
			[]int{1, 2},
			2,
		},
		{
			"multiple duplicates",
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			[]int{0, 1, 2, 3, 4},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.k {
				t.Errorf("want %d, got %d", tt.k, got)
			}
			if diff := cmp.Diff(tt.want, tt.nums[:got]); diff != "" {
				t.Error(diff)
			}
		})
	}
}
