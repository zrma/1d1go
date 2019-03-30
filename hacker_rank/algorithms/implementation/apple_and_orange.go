package data_structures

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	inside := func(offset, distance int32) bool {
		return (s <= offset+distance) && (offset+distance <= t)
	}

	var sum int32
	for _, apple := range apples {
		if inside(a, apple) {
			sum += 1
		}
	}
	fmt.Println(sum)
	sum = 0
	for _, orange := range oranges {
		if inside(b, orange) {
			sum += 1
		}
	}
	fmt.Println(sum)
}
