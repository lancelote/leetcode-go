package p1207

import (
	"testing"
)

func Test_uniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			"true",
			[]int{1, 2, 2, 1, 1, 3},
			true,
		},
		{
			"false",
			[]int{1, 2},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uniqueOccurrences(tt.arr)
			if actual != tt.want {
				t.Errorf("want %t, got %t", tt.want, actual)
			}
		})
	}
}
