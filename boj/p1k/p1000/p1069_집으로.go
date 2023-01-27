package p1000

import (
	"fmt"
	"io"
	"math"

	"1d1go/utils/integer"
)

func Solve1069(reader io.Reader, writer io.Writer) {
	var x, y, d, t int
	_, _ = fmt.Fscan(reader, &x, &y, &d, &t)

	dist := math.Sqrt(float64(x*x + y*y))

	// 한 번도 안 뛰었을 때
	walkTime := dist

	// 최소한 1번은 뛰었을 때
	jumpCnt := integer.Max(int(dist)/d, 1)

	jumAndWalkTime := float64(jumpCnt)*float64(t) + math.Abs(dist-float64(jumpCnt*d))
	jumpOnlyTime := float64(t * (jumpCnt + 1)) // 정확하게 점프로 골인하기 위해 마지막에 점프 한 번 더 함

	_, _ = fmt.Fprintf(writer, "%0.9f", math.Min(walkTime, math.Min(jumAndWalkTime, jumpOnlyTime)))
}
