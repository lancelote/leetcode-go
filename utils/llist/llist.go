package llist

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func ListFromSlice(values []int) *ListNode {
	dummy := newListNode(-1, nil)

	node := dummy
	for _, val := range values {
		next := newListNode(val, nil)
		node.Next = next
		node = next
	}

	return dummy.Next
}

func SliceFromList(node *ListNode) []int {
	slice := []int{}

	for node != nil {
		slice = append(slice, node.Val)
		node = node.Next
	}

	return slice
}
