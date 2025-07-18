package p2352

import (
	"testing"
)

func Test_equalPairs(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"small",
			[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			1,
		},
		{
			"large",
			[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equalPairs(tt.grid)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
