package p392

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			"true case",
			"abc",
			"ahbgdc",
			true,
		},
		{
			"false case",
			"axc",
			"ahbgdc",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isSubsequence(tt.s, tt.t)
			if actual != tt.want {
				t.Error()
			}
		})
	}
}
