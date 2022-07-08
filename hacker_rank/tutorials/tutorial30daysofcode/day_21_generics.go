package tutorial30daysofcode

func printArray(args ...any) {
	for _, arg := range args {
		_, _ = funcPrintln(arg)
	}
}
