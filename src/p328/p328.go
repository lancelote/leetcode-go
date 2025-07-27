package p328

import (
	. "github.com/lancelote/leetcode-go/utils/llist"
)

func oddEvenList(head *ListNode) *ListNode {
	dummyA := &ListNode{}
	dummyB := &ListNode{}

	a, b := dummyA, dummyB

	len := 0
	curNode := head
	for curNode != nil {
		a.Next = curNode
		a = a.Next
		a, b = b, a

		curNode = curNode.Next
		len += 1
	}

	if len%2 == 0 {
		a.Next = dummyB.Next
		b.Next = nil
	} else {
		b.Next = dummyB.Next
		a.Next = nil
	}

	return dummyA.Next
}
