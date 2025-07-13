package btree

import (
	"testing"
)

func Test_newTreeNode(t *testing.T) {
	node := newTreeNode(42)
	if node.Val != 42 {
		t.Errorf("want node value 42, got %d", node.Val)
	}
	if node.Left != nil {
		t.Error("node left is not nil")
	}
	if node.Right != nil {
		t.Error("node right is not nil")
	}
}

func TestNewTreeFromSlice(t *testing.T) {
	slice := []any{1, 2, 3, nil, 4, 5}
	root := NewTreeFromSlice(slice)

	if root.Val != 1 {
		t.Errorf("root val want=%d, got=%d", 1, root.Val)
	}

	if root.Left.Val != 2 {
		t.Errorf("second node val want=%d, got=%d", 2, root.Left.Val)
	}

	if root.Right.Val != 3 {
		t.Errorf("third node val want=%d, got=%d", 3, root.Right.Val)
	}

	if root.Left.Left != nil {
		t.Error("second node is not expected to have a left child")
	}

	if root.Right.Right != nil {
		t.Error("third node is not expected to have a right child")
	}

	if root.Left.Right.Val != 4 {
		t.Errorf("fourth node val want=%d, got=%d", 4, root.Left.Right.Val)
	}

	if root.Right.Left.Val != 5 {
		t.Errorf("fifth node val want=%d, got=%d", 5, root.Right.Left.Val)
	}
}
