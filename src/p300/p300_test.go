package p300

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"first",
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			"second",
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			"same numbers",
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
		{
			"third",
			[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
