package llist

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func testListsEqual(t *testing.T, want *ListNode, got *ListNode) bool {
	for want != nil && got != nil {
		if want.Val != got.Val {
			t.Errorf("want=%d, got=%d", want.Val, got.Val)
			return false
		}

		want = want.Next
		got = got.Next
	}

	if !(want == nil && got == nil) {
		t.Error("unexpected nil")
		return false
	}

	return true
}

func TestListFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  *ListNode
	}{
		{
			"empty",
			[]int{},
			nil,
		},
		{
			"single value",
			[]int{42},
			newListNode(42, nil),
		},
		{
			"long example",
			[]int{1, 2, 3},
			newListNode(1, newListNode(2, newListNode(3, nil))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListFromSlice(tt.slice)
			testListsEqual(t, tt.want, got)
		})
	}
}

func TestSliceFromList(t *testing.T) {
	tests := []struct {
		name string
		node *ListNode
		want []int
	}{
		{
			"nil",
			nil,
			[]int{},
		},
		{
			"single value",
			newListNode(42, nil),
			[]int{42},
		},
		{
			"multipe values",
			newListNode(1, newListNode(2, newListNode(3, nil))),
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceFromList(tt.node)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
