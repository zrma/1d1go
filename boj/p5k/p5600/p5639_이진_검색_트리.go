package p5600

import (
	"fmt"
	"io"
)

func Solve5639(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)
	var root = &node{val: n}
	for {
		_, err := fmt.Fscan(reader, &n)
		if err == io.EOF {
			break
		}
		root.insert(n)
	}
	root.traversePostOrder(writer)
}

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) insert(val int) {
	if val < n.val {
		if n.left == nil {
			n.left = &node{val: val}
		} else {
			n.left.insert(val)
		}
	}

	if val > n.val {
		if n.right == nil {
			n.right = &node{val: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (n *node) traversePostOrder(writer io.Writer) {
	if n == nil {
		return
	}

	n.left.traversePostOrder(writer)
	n.right.traversePostOrder(writer)
	_, _ = fmt.Fprintln(writer, n.val)
}
