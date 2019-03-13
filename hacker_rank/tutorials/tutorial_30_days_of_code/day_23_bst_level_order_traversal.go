package tutorial_30_days_of_code

func levelOrder(root *treeNode) (result []int) {
	queue := make([]*treeNode, 0)

	for root != nil {
		result = append(result, root.data)

		if root.left != nil {
			queue = append(queue, root.left)
		}
		if root.right != nil {
			queue = append(queue, root.right)
		}
		if len(queue) == 0 {
			break
		}

		root = queue[0]
		queue = queue[1:]
	}
	return
}
