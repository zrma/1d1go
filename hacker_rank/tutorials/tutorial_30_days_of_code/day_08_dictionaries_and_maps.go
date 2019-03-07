package tutorial_30_days_of_code

import (
	"fmt"
	"strings"
)

func dictionariesAndMaps(n int, ss []string) {
	m := make(map[string]string)
	for _, s := range ss[:n] {
		tokens := strings.Split(s, " ")
		name, contact := tokens[0], tokens[1]

		m[name] = contact
	}

	for _, s := range ss[n:] {
		if contact, ok := m[s]; ok {
			fmt.Printf("%s=%s\n", s, contact)
		} else {
			fmt.Println("Not found")
		}
	}
}
