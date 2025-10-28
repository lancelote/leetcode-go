package p205

import (
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			"easy true",
			"egg",
			"add",
			true,
		},
		{
			"easy false",
			"foo",
			"bar",
			false,
		},
		{
			"complex true",
			"paper",
			"title",
			true,
		},
		{
			"clash of replacements",
			"badc",
			"baba",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIsomorphic(tt.s, tt.t)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
