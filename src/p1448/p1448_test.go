package p1448

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_goodNodes(t *testing.T) {
	tests := []struct {
		name  string
		slice []any
		want  int
	}{
		{
			"large example",
			[]any{3, 1, 4, 3, nil, 1, 5},
			4,
		},
		{
			"small example",
			[]any{3, 3, nil, 4, 2},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.slice)
			got := goodNodes(root)

			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
