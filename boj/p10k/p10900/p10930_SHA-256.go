package p10900

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func Solve10930(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	sum := sha256.Sum256([]byte(s))
	_, _ = fmt.Fprintf(writer, "%x", sum)
}
