package p643

import (
	"testing"
)

// func findMaxAverage(nums []int, k int) float64 {
func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{
			"basic case",
			[]int{1, 12, -5, -6, 50, 3},
			4,
			12.75,
		},
		{
			"single number",
			[]int{5},
			1,
			5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findMaxAverage(tt.nums, tt.k)
			if actual != tt.want {
				t.Errorf("want %f, got %f", tt.want, actual)
			}
		})
	}
}
