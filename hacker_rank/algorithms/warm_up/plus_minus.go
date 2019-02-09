package warm_up

import "fmt"

func plusMinus(arr []int32) {
	const fmtStr = "%0.6f\n"

	var pos, nag, zero, tot int32
	for _, n := range arr {
		if n > 0 {
			pos++
		}
		if n < 0 {
			nag++
		}
		if n == 0 {
			zero++
		}

		tot++
	}

	fmt.Printf(fmtStr, float64(pos)/float64(tot))
	fmt.Printf(fmtStr, float64(nag)/float64(tot))
	fmt.Printf(fmtStr, float64(zero)/float64(tot))
}

func PlusMinus(arr []int32) {
	plusMinus(arr)
}
