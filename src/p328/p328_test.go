package p328

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/llist"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	tests := []struct {
		name      string
		inSlice   []int
		wantSlice []int
	}{
		{
			"first",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 3, 5, 2, 4},
		},
		{
			"second",
			[]int{2, 1, 3, 5, 6, 4, 7},
			[]int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := llist.ListFromSlice(tt.inSlice)
			gotList := oddEvenList(head)
			gotSlice := llist.SliceFromList(gotList)

			if diff := cmp.Diff(tt.wantSlice, gotSlice); diff != "" {
				t.Error(diff)
			}
		})
	}
}
