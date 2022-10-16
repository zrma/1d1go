package p2200

import (
	"fmt"
)

func Solve2263(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &inorder[i])
	}

	postorder := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &postorder[i])
	}

	inorderIndices := make([]int, n+1)
	for i := 0; i < n; i++ {
		inorderIndices[inorder[i]] = i
	}

	getPreorder(
		inorder, inorderIndices, postorder,
		0, n-1,
		0, n-1,
		writer,
	)
}

func getPreorder(
	inorder, inorderIndices, postorder []int,
	inorderStart int, inorderEnd int,
	postorderStart int, postorderEnd int,
	writer Writer,
) {
	if inorderStart > inorderEnd || postorderStart > postorderEnd {
		return
	}

	root := postorder[postorderEnd]
	_, _ = fmt.Fprintf(writer, "%d ", root)

	left := inorderIndices[root] - inorderStart
	right := inorderEnd - inorderIndices[root]

	getPreorder(
		inorder, inorderIndices, postorder,
		inorderStart, inorderStart+left-1,
		postorderStart, postorderStart+left-1,
		writer,
	)
	getPreorder(
		inorder, inorderIndices, postorder,
		inorderEnd-right+1, inorderEnd,
		postorderEnd-right, postorderEnd-1,
		writer,
	)
}
