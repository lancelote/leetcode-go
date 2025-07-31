package p236

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name    string
		inSlice []any
		p       *btree.TreeNode
		q       *btree.TreeNode
		wantVal int
	}{
		{
			"root is the answer",
			[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			&btree.TreeNode{Val: 5},
			&btree.TreeNode{Val: 1},
			3,
		},
		{
			"p is the answer",
			[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			&btree.TreeNode{Val: 5},
			&btree.TreeNode{Val: 4},
			5,
		},
		{
			"small",
			[]any{1, 2},
			&btree.TreeNode{Val: 1},
			&btree.TreeNode{Val: 2},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.inSlice)
			got := lowestCommonAncestor(root, tt.p, tt.q)
			if got.Val != tt.wantVal {
				t.Errorf("want %d, got %d", tt.wantVal, got.Val)
			}
		})
	}
}
