package p1200

import (
	"fmt"
	"io"
)

func Solve1213(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	n := len(s)
	cnt := make([]int, 26)

	for _, c := range s {
		cnt[c-'A']++
	}

	var odd int
	for _, c := range cnt {
		if c%2 == 1 {
			odd++
		}
	}

	if odd > 1 {
		//goland:noinspection SpellCheckingInspection
		_, _ = fmt.Fprint(writer, "I'm Sorry Hansoo")
		return
	}

	var mid rune
	var ans []rune
	for i, c := range cnt {
		if c%2 == 1 {
			mid = rune(i + 'A')
		}
		for j := 0; j < c/2; j++ {
			ans = append(ans, rune(i+'A'))
		}
	}

	if mid != 0 {
		ans = append(ans, mid)
	}

	for i := n/2 - 1; i >= 0; i-- {
		ans = append(ans, ans[i])
	}

	_, _ = fmt.Fprint(writer, string(ans))
}
