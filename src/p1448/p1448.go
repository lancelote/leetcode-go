package p1448

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	newMax := max(maxVal, node.Val)
	childrenResult := dfs(node.Left, newMax) + dfs(node.Right, newMax)

	if node.Val < maxVal {
		return childrenResult
	}

	return childrenResult + 1
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 1

	if root.Left != nil {
		result += dfs(root.Left, root.Val)
	}

	if root.Right != nil {
		result += dfs(root.Right, root.Val)
	}

	return result
}
