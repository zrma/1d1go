package tutorial_30_days_of_code

import (
	"fmt"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

func moreException(n, p int) error {
	if n < 0 || p < 0 {
		return fmt.Errorf("n and p should be non-negative")
	}

	_, err := fmt.Println(utils.Pow(n, p))
	return err
}

func MoreException(n, p int) {
	if err := moreException(n, p); err != nil {
		fmt.Println(err)
	}
}
