package p1318

import (
	"testing"
)

func Test_minFlips(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{"first", 2, 6, 5, 3},
		{"second", 4, 2, 7, 1},
		{"third", 1, 2, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minFlips(tt.a, tt.b, tt.c)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
