package p136

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"first", []int{2, 2, 1}, 1},
		{"second", []int{4, 1, 2, 1, 2}, 4},
		{"third", []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
