package p135

import (
	"testing"
)

func Test_candy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		{"first", []int{1, 0, 2}, 5},
		{"second", []int{1, 2, 2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy(tt.ratings)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
