package p1137

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"small", 4, 4},
		{"big", 25, 1389537},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tribonacci(tt.n)
			if got != tt.want {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
