package p790

import (
	"testing"
)

func Test_numTiligns(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"first", 3, 5},
		{"second", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numTilings(tt.n)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
