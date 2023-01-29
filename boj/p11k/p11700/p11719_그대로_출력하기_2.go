package p11700

import (
	"bufio"
	"fmt"
	"io"
)

func Solve11719(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, _ = fmt.Fprintln(writer, scanner.Text())
	}
}
