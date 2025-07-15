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

func SliceFromTree(root *TreeNode) []any {
	result := []any{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, node.Val)
		q = append(q, node.Left)
		q = append(q, node.Right)
	}

	// trim nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}
