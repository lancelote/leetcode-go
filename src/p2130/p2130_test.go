package p2130

import (
	"github.com/lancelote/leetcode-go/utils/llist"
	"testing"
)

func Test_pairSum(t *testing.T) {
	tests := []struct {
		name    string
		inSlice []int
		want    int
	}{
		{
			"first",
			[]int{5, 4, 2, 1},
			6,
		},
		{
			"second",
			[]int{4, 2, 2, 3},
			7,
		},
		{
			"third",
			[]int{1, 100000},
			100001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := llist.ListFromSlice(tt.inSlice)
			got := pairSum(head)

			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
