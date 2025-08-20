package p739

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			"first",
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			"second",
			[]int{30, 40, 50, 60},
			[]int{1, 1, 1, 0},
		},
		{
			"third",
			[]int{30, 60, 90},
			[]int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
