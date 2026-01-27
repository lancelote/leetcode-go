package p42

import (
	"testing"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			"first",
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			"second",
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
