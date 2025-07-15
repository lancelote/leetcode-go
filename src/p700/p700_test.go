package p700

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_searchBST(t *testing.T) {
	tests := []struct {
		name  string
		slice []any
		val   int
		want  []any
	}{
		{
			"match",
			[]any{4, 2, 7, 1, 3},
			2,
			[]any{2, 1, 3},
		},
		{
			"none",
			[]any{4, 2, 7, 1, 3},
			5,
			[]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.NewTreeFromSlice(tt.slice)
			gotRoot := searchBST(root, tt.val)
			got := btree.SliceFromTree(gotRoot)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
