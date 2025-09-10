package p1268

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	tests := []struct {
		name       string
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			"first",
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			"second",
			[]string{"havana"},
			"havana",
			[][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := suggestedProducts(tt.products, tt.searchWord)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
