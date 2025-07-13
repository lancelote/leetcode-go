package p872

import (
	. "github.com/lancelote/leetcode-go/utils/btree"
)

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func getLeafs(node *TreeNode, leafs []int) []int {
	if node != nil {
		if isLeaf(node) {
			leafs = append(leafs, node.Val)
		}
		leafs = getLeafs(node.Left, leafs)
		leafs = getLeafs(node.Right, leafs)
	}

	return leafs
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := getLeafs(root1, []int{})
	leafs2 := getLeafs(root2, []int{})

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}
