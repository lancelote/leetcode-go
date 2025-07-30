package p104

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name  string
		slice []any
		want  int
	}{
		{
			"large example",
			[]any{3, 9, 20, nil, nil, 15, 7},
			3,
		},
		{
			"small example",
			[]any{1, nil, 2},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.slice)
			got := maxDepth(root)

			if got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}
