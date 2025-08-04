package p399

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			"basic",
			[][]string{{"a", "b"}, {"b", "d"}, {"d", "c"}},
			[]float64{3, 5, 7},
			[][]string{{"a", "b"}, {"a", "d"}, {"a", "c"}},
			[]float64{3, 15, 105},
		},
		{
			"first",
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2, 3},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6, 0.5, -1, 1, -1},
		},
		{
			"second",
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			"third",
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			[]float64{0.5, 2.0, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation(tt.equations, tt.values, tt.queries)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
