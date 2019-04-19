package tutorial30daysofcode

import "fmt"

func printArray(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
