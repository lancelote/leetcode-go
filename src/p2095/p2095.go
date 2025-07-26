package p2095

import (
	. "github.com/lancelote/leetcode-go/utils/llist"
)

func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	slow, fast := dummy, dummy

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
