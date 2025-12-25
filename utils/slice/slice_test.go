package slice

import (
	"testing"
)

func Test_Diff(t *testing.T) {
	tests := []struct {
		name         string
		want         []int
		got          []int
		expectedDiff string
	}{
		{
			"same empty",
			[]int{},
			[]int{},
			"",
		},
		{
			"same long",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualDiff := Diff(tt.want, tt.got)
			if tt.expectedDiff != actualDiff {
				t.Errorf("want %s, got %s", tt.expectedDiff, actualDiff)
			}
		})
	}
}
