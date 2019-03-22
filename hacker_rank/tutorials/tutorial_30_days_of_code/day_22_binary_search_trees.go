package tutorial_30_days_of_code

import "github.com/zrma/1d1c/hacker_rank/common/utils/integer_util"

type treeNode struct {
	data        int
	left, right *treeNode
}

func insert(root *treeNode, data int) *treeNode {
	if root == nil {
		return &treeNode{data: data}
	}

	var cur *treeNode
	if data <= root.data {
		cur = insert(root.left, data)
		root.left = cur
	} else {
		cur = insert(root.right, data)
		root.right = cur
	}
	return root
}

func getHeight(root *treeNode) int32 {
	if root == nil {
		return -1
	}

	return integer_util.MaxInt32([]int32{getHeight(root.left), getHeight(root.right)}) + 1
}
