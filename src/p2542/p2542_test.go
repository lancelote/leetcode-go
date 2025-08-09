package p2542

import (
	"testing"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{
			"first",
			[]int{1, 3, 3, 2},
			[]int{2, 1, 3, 4},
			3,
			12,
		},
		{
			"second",
			[]int{4, 2, 3, 1, 1},
			[]int{7, 5, 10, 9, 6},
			1,
			30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScore(tt.nums1, tt.nums2, tt.k)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
