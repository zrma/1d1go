package warm_up

import "fmt"

func staircase(n int32) {
	var i, j int32
	for ; i < n; i++ {
		var s string
		for j = n - 1; j >= 0; j-- {
			if i < j {
				s += " "
			} else {
				s += "#"
			}
		}
		fmt.Println(s)
	}
}
