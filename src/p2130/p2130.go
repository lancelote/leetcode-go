package p2130

import (
	. "github.com/lancelote/leetcode-go/utils/llist"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func pairSum(head *ListNode) int {
	len := 0

	node := head
	for node != nil {
		len++
		node = node.Next
	}

	node = head
	twinSums := make(map[int]int)
	for i := 0; i < len/2; i++ {
		twinSums[i] = node.Val
		node = node.Next
	}

	for i := len / 2; i < len; i++ {
		twinSums[len-i-1] += node.Val
		node = node.Next
	}

	maxTwinSum := twinSums[0]
	for i := 0; i < len/2; i++ {
		maxTwinSum = max(maxTwinSum, twinSums[i])
	}

	return maxTwinSum
}
