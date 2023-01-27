package tutorial30daysofcode

func loop(n int32) {
	var i int32
	for i = 1; i <= 10; i++ {
		_, _ = fmtPrintf("%d x %d = %d\n", n, i, n*i)
	}
}
