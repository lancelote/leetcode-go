package p70

import (
	"strconv"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := climbStairs(tt.n)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
