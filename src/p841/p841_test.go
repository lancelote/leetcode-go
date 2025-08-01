package p841

import (
	"testing"
)

func Test_canVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{
			"can visit",
			[][]int{{1}, {2}, {3}, {}},
			true,
		},
		{
			"cannot visit",
			[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canVisitAllRooms(tt.rooms)
			if got != tt.want {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
