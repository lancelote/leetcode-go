package p994

import (
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"basic",
			[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			4,
		},
		{
			"not every can rot",
			[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			-1,
		},
		{
			"already rotted",
			[][]int{{0, 2}},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := orangesRotting(tt.grid)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
