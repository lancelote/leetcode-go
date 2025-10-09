package p55

import (
	"testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			"first",
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			"second",
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
