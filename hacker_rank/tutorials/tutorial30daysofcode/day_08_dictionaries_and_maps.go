package tutorial30daysofcode

import (
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
			_, _ = fmtPrintf("%s=%s\n", s, contact)
		} else {
			_, _ = fmtPrintln("Not found")
		}
	}
}
