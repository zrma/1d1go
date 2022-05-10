package p2400

import (
	"fmt"
)

func Solve2439(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-1-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
