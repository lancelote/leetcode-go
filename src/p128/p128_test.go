package p128

import (
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"first", []int{100, 4, 200, 1, 3, 2}, 4},
		{"second", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"third", []int{1, 0, 1, 2}, 3},
		{"single element", []int{1}, 1},
		{"empty", []int{}, 0},
		{"two zeros", []int{0, 0}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
