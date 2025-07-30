package p1372

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestZigZag(root *TreeNode) int {
	longestPath := 0

	var dfs func(node *TreeNode, left, right int)
	dfs = func(node *TreeNode, left int, right int) {
		if node == nil {
			return
		}

		dfs(node.Left, right+1, 0)
		dfs(node.Right, 0, left+1)

		longestPath = max(longestPath, left)
		longestPath = max(longestPath, right)
	}

	dfs(root.Left, 1, 0)
	dfs(root.Right, 0, 1)

	return longestPath
}
