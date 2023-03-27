package p15900

import (
	"bufio"
	"fmt"
	"io"
)

func Solve15904(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()

	//goland:noinspection SpellCheckingInspection
	const ucpc = "UCPC"

	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ucpc[idx] {
			idx++
			if idx == len(ucpc) {
				break
			}
		}
	}

	if idx == len(ucpc) {
		//goland:noinspection SpellCheckingInspection
		_, _ = fmt.Fprint(writer, "I love UCPC")
	} else {
		//goland:noinspection SpellCheckingInspection
		_, _ = fmt.Fprint(writer, "I hate UCPC")
	}
}
