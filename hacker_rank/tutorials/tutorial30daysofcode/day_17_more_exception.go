package tutorial30daysofcode

import (
	"fmt"

	"1d1go/utils/integer"
)

func exception(n, p int64) error {
	if n < 0 || p < 0 {
		return fmt.Errorf("n and p should be non-negative")
	}

	_, err := fmtPrint(integer.Pow(n, p))
	return err
}

func moreException(n, p int64) {
	if err := exception(n, p); err != nil {
		_, _ = fmtPrint(err)
	}
}
