package p10100

import (
	"fmt"
	"io"
)

func Solve10171(writer io.Writer) {
	const s = `\    /\
 )  ( ')
(  /  )
 \(__)|`

	_, _ = fmt.Fprint(writer, s)
}
