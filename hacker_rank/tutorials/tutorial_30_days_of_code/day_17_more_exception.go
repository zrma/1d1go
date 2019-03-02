package tutorial_30_days_of_code

import (
	"fmt"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

func moreException(n, p int64) error {
	if n < 0 || p < 0 {
		return fmt.Errorf("n and p should be non-negative")
	}

	_, err := fmt.Println(utils.PowInt64(n, p))
	return err
}

func MoreException(n, p int64) {
	if err := moreException(n, p); err != nil {
		fmt.Println(err)
	}
}
