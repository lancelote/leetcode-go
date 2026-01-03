package p134

import (
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			"has solution",
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			"no solution",
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
