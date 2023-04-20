package p20200

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func Solve20291(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	extensions := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		_, _ = fmt.Fscanln(reader, &name)

		ext := name[strings.LastIndex(name, ".")+1:]
		extensions[ext]++
	}

	var keys []string
	for k := range extensions {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		_, _ = fmt.Fprintf(writer, "%s %d\n", k, extensions[k])
	}
}
