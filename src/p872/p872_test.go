package p872

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []any
		slice2 []any
		want   bool
	}{
		{
			"true case",
			[]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4},
			[]any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8},
			true,
		},
		{
			"false case",
			[]any{1, 2, 3},
			[]any{1, 3, 2},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root1 := btree.NewTreeFromSlice(tt.slice1)
			root2 := btree.NewTreeFromSlice(tt.slice2)

			got := leafSimilar(root1, root2)
			if got != tt.want {
				t.Error()
			}
		})
	}
}
