package p1700

import (
	"fmt"
	"io"
	"sort"
)

func Solve1755(reader io.Reader, writer io.Writer) {
	var m, n int
	_, _ = fmt.Fscanln(reader, &m, &n)

	var words []struct {
		num  int
		word string
	}
	for i := m; i <= n; i++ {
		words = append(words, struct {
			num  int
			word string
		}{i, numToWord(i)})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].word < words[j].word
	})

	for i, w := range words {
		_, _ = fmt.Fprint(writer, w.num)
		if (i+1)%10 == 0 {
			_, _ = fmt.Fprintln(writer)
		} else if i != len(words)-1 {
			_, _ = fmt.Fprint(writer, " ")
		}
	}
}

func numToWord(i int) string {
	var res string
	s := fmt.Sprintf("%d", i)
	for i, c := range s {
		switch c {
		case '1':
			res += "one"
		case '2':
			res += "two"
		case '3':
			res += "three"
		case '4':
			res += "four"
		case '5':
			res += "five"
		case '6':
			res += "six"
		case '7':
			res += "seven"
		case '8':
			res += "eight"
		case '9':
			res += "nine"
		case '0':
			res += "zero"
		}
		if i != len(s)-1 {
			res += " "
		}
	}
	return res
}
