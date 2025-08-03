package p1466

import (
	"testing"
)

func Test_minReorder(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        int
	}{
		{
			"first",
			6,
			[][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			3,
		},
		{
			"second",
			5,
			[][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			2,
		},
		{
			"third",
			3,
			[][]int{{1, 0}, {2, 0}},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minReorder(tt.n, tt.connections)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
