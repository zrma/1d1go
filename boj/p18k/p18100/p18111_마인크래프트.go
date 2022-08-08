package p18100

import (
	"fmt"
	"math"
)

func Solve18111(reader Reader, writer Writer) {
	var n, m, b int
	_, _ = fmt.Fscan(reader, &n, &m, &b)

	tot := b
	const maxSize = 256
	heightCounts := make([]int, maxSize+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var v int
			_, _ = fmt.Fscan(reader, &v)
			tot += v
			heightCounts[v]++
		}
	}

	var minTime = math.MaxInt32
	var maxHeight = 0
	for targetHeight := 0; targetHeight <= maxSize; targetHeight++ {
		if tot < n*m*targetHeight {
			break
		}
		var time int
		for height, count := range heightCounts {
			switch {
			case height < targetHeight:
				time += count * (targetHeight - height) // add step.2 +1
			case height > targetHeight:
				time += count * (height - targetHeight) * 2 // remove step.1 -2
			}
		}
		if time <= minTime {
			minTime = time
			maxHeight = targetHeight
		}
	}
	_, _ = fmt.Fprint(writer, minTime, maxHeight)
}
