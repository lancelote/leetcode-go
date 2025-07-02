package p1679

import (
	"testing"
)

func Test_maxOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			"consume all",
			[]int{1, 2, 3, 4},
			5,
			2,
		},
		{
			"consume once",
			[]int{3, 1, 3, 4, 3},
			6,
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxOperations(tt.nums, tt.k)
			if actual != tt.want {
				t.Errorf("want %d, got %d", tt.want, actual)
			}
		})
	}
}
