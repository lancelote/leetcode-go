package p1431

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			"first",
			[]int{2, 3, 5, 1, 3},
			3,
			[]bool{true, true, true, false, true},
		},
		{
			"second",
			[]int{4, 2, 1, 1, 2},
			1,
			[]bool{true, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kidsWithCandies(tt.candies, tt.extraCandies)
			if diff := cmp.Diff(tt.want, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}
