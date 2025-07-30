package p206

import (
	. "github.com/lancelote/leetcode-go/utils/llist"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}

	return prev
}
