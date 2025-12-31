package p322

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			"basic examples",
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			"impossible",
			[]int{2},
			3,
			-1,
		},
		{
			"zero",
			[]int{1},
			0,
			0,
		},
		{
			"long computation",
			[]int{1, 2, 5},
			100,
			20,
		},
		{
			"very long computation",
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24},
			9999,
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
