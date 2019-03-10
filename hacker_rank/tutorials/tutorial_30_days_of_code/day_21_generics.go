package tutorial_30_days_of_code

import "fmt"

func printArray(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
