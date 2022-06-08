package tutorial30daysofcode

func printArray(args ...interface{}) {
	for _, arg := range args {
		_, _ = funcPrintln(arg)
	}
}
