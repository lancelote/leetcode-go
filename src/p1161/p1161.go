package p1161

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLvlIdx := 1
	maxLvlSum := root.Val

	lvl := []*TreeNode{root}
	lvlIdx := 1

	for len(lvl) > 0 {
		newLvl := []*TreeNode{}
		lvlSum := 0

		for _, node := range lvl {
			lvlSum += node.Val

			if node.Left != nil {
				newLvl = append(newLvl, node.Left)
			}

			if node.Right != nil {
				newLvl = append(newLvl, node.Right)
			}
		}

		if lvlSum > maxLvlSum {
			maxLvlSum = lvlSum
			maxLvlIdx = lvlIdx
		}

		lvlIdx++
		lvl = newLvl
	}

	return maxLvlIdx
}
