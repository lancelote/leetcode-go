package p735

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{
			"two remain",
			[]int{5, 10, -5},
			[]int{5, 10},
		},
		{
			"no remain",
			[]int{8, -8},
			[]int{},
		},
		{
			"one remain",
			[]int{10, 2, -5},
			[]int{10},
		},
		{
			"none will clash",
			[]int{-2, -1, 1, 2},
			[]int{-2, -1, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
