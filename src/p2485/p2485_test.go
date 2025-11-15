package p2485

import (
	"testing"
)

func Test_pivotInteger(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"base", 8, 6},
		{"single", 1, 1},
		{"none", 4, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotInteger(tt.n)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
