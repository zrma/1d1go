package p4700

import (
	"fmt"
	"io"
	"sort"
)

func Solve4716(reader io.Reader, writer io.Writer) {
	for {
		var n, a, b int
		_, _ = fmt.Fscan(reader, &n, &a, &b)
		if n == 0 {
			break
		}

		type balloon struct {
			k, distA, distB int
		}
		balloons := make([]balloon, n)
		for i := range balloons {
			_, _ = fmt.Fscan(reader, &balloons[i].k, &balloons[i].distA, &balloons[i].distB)
		}

		diff := func(v balloon) int {
			dist := v.distA - v.distB
			if dist < 0 {
				return -dist
			}
			return dist
		}

		sort.Slice(balloons, func(i, j int) bool {
			return diff(balloons[i]) > diff(balloons[j])
		})

		var totDist int
		for _, v := range balloons {
			if v.distA < v.distB {
				if a < v.k {
					v.k -= a
					totDist += a * v.distA
					a = 0

					totDist += v.k * v.distB
					b -= v.k
				} else {
					a -= v.k
					totDist += v.k * v.distA
				}
			} else {
				if b < v.k {
					v.k -= b
					totDist += b * v.distB
					b = 0

					totDist += v.k * v.distA
					a -= v.k
				} else {
					b -= v.k
					totDist += v.k * v.distB
				}
			}
		}

		_, _ = fmt.Fprintln(writer, totDist)
	}
}
