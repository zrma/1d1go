package tutorial30daysofcode

import (
	"fmt"
	"github.com/zrma/1d1c/hacker_rank/common/utils/integer"
)

func exception(n, p int64) error {
	if n < 0 || p < 0 {
		return fmt.Errorf("n and p should be non-negative")
	}

	_, err := fmt.Println(integer.PowInt64(n, p))
	return err
}

func moreException(n, p int64) {
	if err := exception(n, p); err != nil {
		fmt.Println(err)
	}
}
