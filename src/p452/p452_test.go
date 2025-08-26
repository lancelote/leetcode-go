package p452

import (
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{"first", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{"second", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{"third", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinArrowShots(tt.points)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
