package p10100

import (
	"fmt"
	"io"
)

func Solve10172(writer io.Writer) {
	const s = "|\\_/|\n|q p|   /}\n( 0 )\"\"\"\\\n|\"^\"`    |\n||_/=\\\\__|"

	_, _ = fmt.Fprint(writer, s)
}
