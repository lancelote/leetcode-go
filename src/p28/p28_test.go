package p28

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "ll", 2},
		{"a", "a", 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.haystack, tt.needle), func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
