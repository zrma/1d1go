package ps_basic

import (
	"sort"
)

func mostActive(customers []string) []string {
	var en entries
	for _, c := range customers {
		en = en.increase(c)
	}

	sort.Slice(en, func(i, j int) bool {
		return en[i].count > en[j].count
	})
	sum := 0
	for _, e := range en {
		sum += e.count
	}
	var res []string
	for _, e := range en {
		if float64(e.count)/float64(sum) >= 0.05 {
			res = append(res, e.name)
		}
	}
	sort.Strings(res)
	return res
}

type entry struct {
	name  string
	count int
}

type entries []entry

func (en entries) increase(customer string) entries {
	for i, e := range en {
		if e.name == customer {
			en[i].count++
			return en
		}
	}
	return append(en, entry{customer, 1})
}
