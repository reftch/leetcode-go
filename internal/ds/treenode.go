package ds

// TreeNode is the shared binary tree node used by LeetCode problems.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode builds a binary tree from a level-order slice where nil means
// a missing node, e.g. []any{1, 2, 3, nil, 4}.
func NewTreeNode(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	toInt := func(v any) int {
		switch n := v.(type) {
		case int:
			return n
		case int8:
			return int(n)
		case int16:
			return int(n)
		case int32:
			return int(n)
		case int64:
			return int(n)
		case uint:
			return int(n)
		case float64:
			return int(n)
		default:
			return 0
		}
	}
	root := &TreeNode{Val: toInt(vals[0])}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: toInt(vals[i])}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: toInt(vals[i])}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// ToLevelOrder converts a tree back to level-order form, trimming trailing nils.
func (n *TreeNode) ToLevelOrder() []any {
	if n == nil {
		return nil
	}
	var out []any
	queue := []*TreeNode{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}
