package p11

import (
	"testing"
)

// func maxArea(height []int) int {
func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			"large example",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"small example",
			[]int{1, 1},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxArea(tt.height)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
