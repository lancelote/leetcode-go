package p1502

import (
	"testing"
)

func Test_canMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			"true",
			[]int{3, 5, 1},
			true,
		},
		{
			"false",
			[]int{1, 2, 4},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canMakeArithmeticProgression(tt.arr)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
