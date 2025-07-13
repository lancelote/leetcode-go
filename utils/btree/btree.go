package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func NewTreeFromSlice(values []any) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := newTreeNode(values[0].(int))
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			left := newTreeNode(values[i].(int))
			queue = append(queue, left)
			node.Left = left
		}
		i++

		if i < len(values) && values[i] != nil {
			right := newTreeNode(values[i].(int))
			queue = append(queue, right)
			node.Right = right
		}
		i++
	}

	return root
}
