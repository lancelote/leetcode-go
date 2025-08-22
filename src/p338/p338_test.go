package p338

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"first", 2, []int{0, 1, 1}},
		{"second", 5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
