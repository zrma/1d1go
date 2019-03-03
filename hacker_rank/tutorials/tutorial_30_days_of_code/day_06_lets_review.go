package tutorial_30_days_of_code

import "fmt"

func letsReview(s string) {
	var even, odd string
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			even += string(s[i])
		} else {
			odd += string(s[i])
		}
	}
	fmt.Println(even, odd)
}

func LetsReview(ss []string) {
	for _, s := range ss {
		letsReview(s)
	}
}
