package p437

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	cache := make(map[int]int)

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}

		curSum += node.Val

		if curSum == targetSum {
			count += 1
		}

		if v, ok := cache[curSum-targetSum]; ok {
			count += v
		}

		cache[curSum]++

		dfs(node.Left, curSum)
		dfs(node.Right, curSum)

		cache[curSum]--
	}

	dfs(root, 0)
	return count
}
