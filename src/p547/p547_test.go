package p547

import (
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			"two groups",
			[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			2,
		},
		{
			"no connections",
			[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCircleNum(tt.isConnected)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
