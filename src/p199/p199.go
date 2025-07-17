package p199

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Val}
	layer := []*TreeNode{root}

	for len(layer) > 0 {
		newLayer := []*TreeNode{}

		for _, node := range layer {
			if node.Left != nil {
				newLayer = append(newLayer, node.Left)
			}

			if node.Right != nil {
				newLayer = append(newLayer, node.Right)
			}
		}

		if len(newLayer) > 0 {
			result = append(result, newLayer[len(newLayer)-1].Val)
		}

		layer = newLayer
	}

	return result
}
