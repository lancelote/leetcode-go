package p169

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"first",
			[]int{3, 2, 3},
			3,
		},
		{
			"second",
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
