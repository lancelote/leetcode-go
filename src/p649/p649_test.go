package p649

import (
	"testing"
)

func Test_predictPartyVictory(t *testing.T) {
	tests := []struct {
		name   string
		senate string
		want   string
	}{
		{
			"R win",
			"RD",
			"Radiant",
		},
		{
			"D win",
			"RDD",
			"Dire",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := predictPartyVictory(tt.senate)
			if got != tt.want {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
