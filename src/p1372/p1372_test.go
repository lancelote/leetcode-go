package p1372

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_longestZigZag(t *testing.T) {
	tests := []struct {
		name    string
		inSlice []any
		want    int
	}{
		{
			"first",
			[]any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1},
			3,
		},
		{
			"second",
			[]any{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1},
			4,
		},
		{
			"third",
			[]any{1},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.inSlice)
			got := longestZigZag(root)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
