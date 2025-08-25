package p435

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			"remove one",
			[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			1,
		},
		{
			"remove two",
			[][]int{{1, 2}, {1, 2}, {1, 2}},
			2,
		},
		{
			"nothing to remove",
			[][]int{{1, 2}, {2, 3}},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eraseOverlapIntervals(tt.intervals)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
