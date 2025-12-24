package p3074

import (
	"testing"
)

func Test_minimumBoxes(t *testing.T) {
	tests := []struct {
		name     string
		apple    []int
		capacity []int
		want     int
	}{
		{
			"first",
			[]int{1, 3, 2},
			[]int{4, 3, 1, 5, 2},
			2,
		},
		{
			"second",
			[]int{5, 5, 5},
			[]int{2, 4, 2, 7},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumBoxes(tt.apple, tt.capacity)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
