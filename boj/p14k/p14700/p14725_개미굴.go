package p14700

import (
	"fmt"
	"io"
	"sort"
)

func Solve14725(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	var root = &trie{}

	for i := 0; i < n; i++ {
		var k int
		_, _ = fmt.Fscan(reader, &k)
		var cur = root
		for j := 0; j < k; j++ {
			var s string
			_, _ = fmt.Fscan(reader, &s)
			cur = cur.add(s)
		}
	}

	root.print(writer, 0)
}

type trie struct {
	children map[string]*trie
}

func (t *trie) add(s string) *trie {
	if t.children == nil {
		t.children = make(map[string]*trie)
	}
	if t.children[s] == nil {
		t.children[s] = &trie{}
	}
	return t.children[s]
}

func (t *trie) print(writer io.Writer, depth int) {
	if t.children == nil {
		return
	}

	keys := make([]string, 0, len(t.children))
	for k := range t.children {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		v := t.children[k]
		for i := 0; i < depth; i++ {
			_, _ = fmt.Fprint(writer, "--")
		}
		_, _ = fmt.Fprintln(writer, k)
		v.print(writer, depth+1)
	}
}
