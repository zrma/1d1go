package p5600

import (
	"fmt"
	"io"
	"sort"
	"time"
)

func Solve5635(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	type person struct {
		name     string
		birthday time.Time
	}

	persons := make([]person, n)
	for i := 0; i < n; i++ {
		var name string
		var day, month, year int
		_, _ = fmt.Fscanln(reader, &name, &day, &month, &year)
		persons[i] = person{name, time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].birthday.After(persons[j].birthday)
	})

	_, _ = fmt.Fprintln(writer, persons[0].name)
	_, _ = fmt.Fprint(writer, persons[n-1].name)
}
