package p49

import (
	"github.com/google/go-cmp/cmp"
	"slices"
	"strings"
	"testing"
)

func sortResult(r [][]string) {
	for _, group := range r {
		slices.Sort(group)
	}

	slices.SortFunc(r, func(a, b []string) int {
		return strings.Compare(a[0], b[0])
	})
}

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			"base case",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			"empty string",
			[]string{""},
			[][]string{{""}},
		},
		{
			"single letter",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			sortResult(got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
