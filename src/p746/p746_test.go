package p746

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			"first",
			[]int{10, 15, 20},
			15,
		},
		{
			"second",
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
