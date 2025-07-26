package p2095

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/llist"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	tests := []struct {
		name      string
		inSlice   []int
		wantSlice []int
	}{
		{
			"first",
			[]int{1, 3, 4, 7, 1, 2, 6},
			[]int{1, 3, 4, 1, 2, 6},
		},
		{
			"second",
			[]int{1, 2, 3, 4},
			[]int{1, 2, 4},
		},
		{
			"third",
			[]int{2, 1},
			[]int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := llist.ListFromSlice(tt.inSlice)
			got := deleteMiddle(head)
			gotSlice := llist.SliceFromList(got)

			if diff := cmp.Diff(tt.wantSlice, gotSlice); diff != "" {
				t.Error(diff)
			}
		})
	}
}
