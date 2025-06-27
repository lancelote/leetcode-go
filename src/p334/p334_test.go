package p334

import (
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			"sorted asc",
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			"sorted desc",
			[]int{5, 4, 3, 2, 1},
			false,
		},
		{
			"mix",
			[]int{2, 1, 5, 0, 4, 6},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := increasingTriplet(tt.nums)
			if actual != tt.want {
				t.Error()
			}
		})
	}
}
