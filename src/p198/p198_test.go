package p198

import (
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"first",
			[]int{1, 2, 3, 1},
			4,
		},
		{
			"second",
			[]int{2, 7, 9, 3, 1},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
