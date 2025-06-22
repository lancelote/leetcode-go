package p605

import (
	"testing"
)

// func canPlaceFlowers(flowerbed []int, n int) bool {
func Test_canPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			"can plant one",
			[]int{1, 0, 0, 0, 1},
			1,
			true,
		},
		{
			"cann't plant two",
			[]int{1, 0, 0, 0, 1},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canPlaceFlowers(tt.flowerbed, tt.n)
			if result != tt.want {
				t.Error()
			}
		})
	}
}
