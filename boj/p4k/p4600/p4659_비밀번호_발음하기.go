package p4600

import (
	"fmt"
	"io"
)

func Solve4659(reader io.Reader, writer io.Writer) {
	for {
		var s string
		_, _ = fmt.Fscanln(reader, &s)
		if s == "end" {
			break
		}

		isAcceptable := true
		var vowel, consonant, vowelCount, consonantCount int
		var prev byte

		for i := 0; i < len(s); i++ {
			switch s[i] {
			case 'a', 'e', 'i', 'o', 'u':
				vowel++
				consonant = 0
				vowelCount++
			default:
				consonant++
				vowel = 0
				consonantCount++
			}
			if vowel == 3 || consonant == 3 {
				isAcceptable = false
				break
			}
			if prev == s[i] && prev != 'e' && prev != 'o' {
				isAcceptable = false
				break
			}
			prev = s[i]
		}

		if vowelCount == 0 {
			isAcceptable = false
		}

		if isAcceptable {
			_, _ = fmt.Fprintln(writer, "<"+s+"> is acceptable.")
		} else {
			_, _ = fmt.Fprintln(writer, "<"+s+"> is not acceptable.")
		}
	}
}
