package p2462

import (
	"testing"
)

func Test_totalCost(t *testing.T) {
	tests := []struct {
		name       string
		costs      []int
		k          int
		candidates int
		want       int64
	}{
		{
			name:       "first",
			costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
			k:          3,
			candidates: 4,
			want:       11,
		},
		{
			name:       "second",
			costs:      []int{1, 2, 4, 1},
			k:          3,
			candidates: 3,
			want:       4,
		},
		{
			name:       "third",
			costs:      []int{17, 12, 2, 200, 8},
			k:          1,
			candidates: 3,
			want:       2,
		},
		{
			name:       "forth",
			costs:      []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58},
			k:          11,
			candidates: 2,
			want:       423,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalCost(tt.costs, tt.k, tt.candidates)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
