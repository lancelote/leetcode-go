package p274

import (
	"testing"
)

func Test_hIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{
			"first",
			[]int{3, 0, 6, 1, 5},
			3,
		},
		{
			"second",
			[]int{1, 3, 1},
			1,
		},
		{
			"third",
			[]int{100},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hIndex(tt.citations)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
