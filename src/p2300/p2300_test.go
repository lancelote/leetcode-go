package p2300

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	tests := []struct {
		name    string
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{
			"first",
			[]int{5, 1, 3},
			[]int{1, 2, 3, 4, 5},
			7,
			[]int{4, 0, 3},
		},
		{
			"second",
			[]int{3, 1, 2},
			[]int{8, 5, 8},
			16,
			[]int{2, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := successfulPairs(tt.spells, tt.potions, tt.success)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
