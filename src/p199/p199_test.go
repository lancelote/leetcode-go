package p199

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name  string
		slice []any
		want  []int
	}{
		{
			"first",
			[]any{1, 2, 3, nil, 5, nil, 4},
			[]int{1, 3, 4},
		},
		{
			"second",
			[]any{1, 2, 3, 4, nil, nil, nil, 5},
			[]int{1, 3, 4, 5},
		},
		{
			"small",
			[]any{1, nil, 3},
			[]int{1, 3},
		},
		{
			"empty",
			[]any{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.NewTreeFromSlice(tt.slice)
			got := rightSideView(root)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
