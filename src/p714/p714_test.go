package p714

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		fee    int
		want   int
	}{
		{
			"first",
			[]int{1, 3, 2, 8, 4, 9},
			2,
			8,
		},
		{
			"second",
			[]int{1, 3, 7, 5, 10, 3},
			3,
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices, tt.fee)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
