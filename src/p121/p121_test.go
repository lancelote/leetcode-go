package p121

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			"basic case",
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			"no profit",
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
