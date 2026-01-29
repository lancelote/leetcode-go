package p12

import (
	"strconv"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		name := strconv.Itoa(tt.num)
		t.Run(name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
