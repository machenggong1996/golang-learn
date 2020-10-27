package tree

func preorderTraversal(root *TreeNode) (vals []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		//获取最后一个元素的右节点
		node = stack[len(stack)-1].Right
		//将最后一个元素取出
		stack = stack[:len(stack)-1]
	}
	return
}
