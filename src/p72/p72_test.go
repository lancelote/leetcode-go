package p72

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			"first",
			"horse",
			"ros",
			3,
		},
		{
			"second",
			"intention",
			"execution",
			5,
		},
		{
			"third",
			"a",
			"aa",
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
