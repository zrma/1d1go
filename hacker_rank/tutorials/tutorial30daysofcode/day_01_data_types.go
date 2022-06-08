package tutorial30daysofcode

func dataType(i1, i2 uint64, d1, d2 float64, s1, s2 string) {
	_, _ = funcPrintln(i1 + i2)
	_, _ = funcPrintf("%.1f\n", d1+d2)
	_, _ = funcPrintln(s1 + s2)
}
