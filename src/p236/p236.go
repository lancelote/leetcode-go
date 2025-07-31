package p236

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := root

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		mid := node.Val == p.Val || node.Val == q.Val

		if left && right || mid && left || mid && right {
			result = node
		}

		return left || mid || right
	}

	dfs(root)
	return result
}
