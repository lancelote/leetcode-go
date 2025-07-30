package p206

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lancelote/leetcode-go/utils/llist"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name      string
		inSlice   []int
		wantSlice []int
	}{
		{
			"large",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"small",
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"empty",
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := llist.ListFromSlice(tt.inSlice)
			gotList := reverseList(head)
			gotSlice := llist.SliceFromList(gotList)

			if diff := cmp.Diff(tt.wantSlice, gotSlice); diff != "" {
				t.Error(diff)
			}
		})
	}
}
