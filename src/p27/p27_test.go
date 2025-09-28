package p27

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		k    int
		want []int
	}{
		{
			"first",
			[]int{3, 2, 2, 3},
			3,
			2,
			[]int{2, 2},
		},
		{
			"second",
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
			[]int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if tt.k != got {
				t.Errorf("want %d, got %d", tt.k, got)
			}
			if diff := cmp.Diff(tt.want, tt.nums[:got]); diff != "" {
				t.Error(diff)
			}
		})
	}
}
