package p437

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_pathSum(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		targetSum int
		want      int
	}{
		{
			"first",
			[]any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1},
			8,
			3,
		},
		{
			"second",
			[]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
			22,
			3,
		},
		{
			"small",
			[]any{1, 2},
			0,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.slice)
			got := pathSum(root, tt.targetSum)

			if got != tt.want {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
