package p1900

import (
	"fmt"
	"io"
)

func Solve1991(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]node, 26)
	for i := 0; i < n; i++ {
		var val, left, right string
		_, _ = fmt.Fscan(reader, &val, &left, &right)

		idx := val[0] - 'A'
		arr[idx] = node{left, right}
	}

	traversePreOrder("A", arr, writer)
	_, _ = fmt.Fprintln(writer)

	traverseInOrder("A", arr, writer)
	_, _ = fmt.Fprintln(writer)

	traversePostOrder("A", arr, writer)
	_, _ = fmt.Fprintln(writer)
}

type node struct {
	left  string
	right string
}

func traversePreOrder(root string, arr []node, writer io.Writer) {
	if root == "." {
		return
	}
	idx := root[0] - 'A'
	_, _ = fmt.Fprint(writer, root)
	traversePreOrder(arr[idx].left, arr, writer)
	traversePreOrder(arr[idx].right, arr, writer)
}

func traverseInOrder(root string, arr []node, writer io.Writer) {
	if root == "." {
		return
	}
	idx := root[0] - 'A'
	traverseInOrder(arr[idx].left, arr, writer)
	_, _ = fmt.Fprint(writer, root)
	traverseInOrder(arr[idx].right, arr, writer)
}

func traversePostOrder(root string, arr []node, writer io.Writer) {
	if root == "." {
		return
	}
	idx := root[0] - 'A'
	traversePostOrder(arr[idx].left, arr, writer)
	traversePostOrder(arr[idx].right, arr, writer)
	_, _ = fmt.Fprint(writer, root)
}
