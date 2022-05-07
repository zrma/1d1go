package codesprint2019

func alphaBeta(arr []int32) int32 {
	var alpha, beta int32
	for _, cur := range arr {
		if cur > alpha {
			alpha, beta = cur, alpha
		} else if cur < alpha && cur > beta {
			beta = cur
		}
	}
	return beta
}
