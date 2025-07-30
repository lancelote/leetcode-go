package p1161

import (
	"github.com/lancelote/leetcode-go/utils/btree"
	"testing"
)

func Test_maxLevelSum(t *testing.T) {
	tests := []struct {
		name  string
		slice []any
		want  int
	}{
		{
			"first",
			[]any{1, 7, 0, 7, -8, nil, nil},
			2,
		},
		{
			"second",
			[]any{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.TreeFromSlice(tt.slice)
			got := maxLevelSum(root)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
