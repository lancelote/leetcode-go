package p162

import (
	"testing"
)

func Test_findPeakElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"first",
			[]int{1, 2, 3, 1},
			2,
		},
		{
			"second",
			[]int{1, 2, 1, 3, 5, 6, 4},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
