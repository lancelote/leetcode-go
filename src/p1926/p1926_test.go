package p1926

import (
	"testing"
)

func Test_nearestExit(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		want     int
	}{
		{
			"first",
			[][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			[]int{1, 2},
			1,
		},
		{
			"second",
			[][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			[]int{1, 0},
			2,
		},
		{
			"third",
			[][]byte{{'.', '+'}},
			[]int{0, 0},
			-1,
		},
		{
			"fourth",
			[][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			[]int{1, 2},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nearestExit(tt.maze, tt.entrance)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
