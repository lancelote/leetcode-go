package p45

import (
	"testing"
)

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"first",
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			"second",
			[]int{2, 3, 0, 1, 4},
			2,
		},
		{
			"single element",
			[]int{0},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
