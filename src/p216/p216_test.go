package p216

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{"single answer", 3, 7, [][]int{{1, 2, 4}}},
		{"multiple answers", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"no answer", 4, 1, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
