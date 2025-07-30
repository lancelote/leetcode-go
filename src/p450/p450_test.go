package p450

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		name      string
		inSlice   []any
		key       int
		wantSlice []any
	}{
		{
			"key exist",
			[]any{5, 3, 6, 2, 4, nil, 7},
			3,
			[]any{5, 4, 6, 2, nil, nil, 7},
		},
		{
			"key doesn't exit",
			[]any{5, 3, 6, 2, 4, nil, 7},
			0,
			[]any{5, 3, 6, 2, 4, nil, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.inSlice)
			got := deleteNode(root, tt.key)
			gotSlice := btree.SliceFromTree(got)

			if diff := cmp.Diff(tt.wantSlice, gotSlice); diff != "" {
				t.Error(diff)
			}
		})
	}
}
