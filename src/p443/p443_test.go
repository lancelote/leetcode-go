package p443

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		name      string
		chars     []byte
		wantInt   int
		wantSlice []byte
	}{
		{
			"basic case",
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			6,
			[]byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			"single char",
			[]byte{'a'},
			1,
			[]byte{'a'},
		},
		{
			"big num",
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			4,
			[]byte{'a', 'b', '1', '2'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualInt := compress(tt.chars)
			if tt.wantInt != actualInt {
				t.Errorf("want %d, got %d", tt.wantInt, actualInt)
			}
			if diff := cmp.Diff(tt.wantSlice, tt.chars[:actualInt]); diff != "" {
				t.Error(diff)
			}
		})
	}
}
