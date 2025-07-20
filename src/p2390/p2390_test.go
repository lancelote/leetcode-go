package p2390

import (
	"testing"
)

func Test_removeStars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"parts",
			"leet**cod*e",
			"lecoe",
		},
		{
			"erase complete",
			"erase*****",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeStars(tt.s)
			if got != tt.want {
				t.Errorf("want=%s, got=%s", tt.want, got)
			}
		})
	}
}
