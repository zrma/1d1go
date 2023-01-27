package p25000

import (
	"fmt"
	"io"
)

func Solve25083(writer io.Writer) {
	const msg = "         ,r'\"7\nr`-_   ,'  ,/\n \\. \". L_r'\n   `~\\/\n      |\n      |"
	_, _ = fmt.Fprint(writer, msg)
}
