package dictionaries_and_hashmaps

import (
	"fmt"
	"sort"
	"strings"
)

type adapter []string

func (a adapter) Len() int           { return len(a) }
func (a adapter) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }
func (a adapter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func checkMagazine(magazine []string, note []string) {
	sort.Sort(adapter(magazine))
	sort.Sort(adapter(note))

	i := len(magazine) - 1
	j := len(note) - 1

	for i >= 0 && j >= 0 {
		diff := strings.Compare(magazine[i], note[j])
		if diff > 0 {
			i--
		} else if diff == 0 {
			i--
			j--
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func CheckMagazine(magazine []string, note []string) {
	checkMagazine(magazine, note)
}
