package p1700

import (
	"fmt"
	"io"
)

func Solve1700(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, k+1)
	for i := 1; i <= k; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	unplugCnt := 0
	pluggedCnt := 0
	plugged := make([]bool, k+1)
	for i := 1; i <= k; i++ {
		curr := arr[i]

		// 꽂혀 있으면 통과
		if plugged[curr] {
			continue
		}

		// 슬롯이 남아 있으면 꽂고 통과
		if pluggedCnt < n {
			plugged[curr] = true
			pluggedCnt++
			continue
		}

		// 슬롯이 부족하므로 빼야 함

		// 각 값들이 이후 등장 할 위치 기록용 슬라이스 초기화
		appearIndices := make([]int, k+1)
		for v := 1; v <= k; v++ {
			appearIndices[v] = -1
		}

		// 각 값들이 이후 등장할 위치
		for idx := i + 1; idx <= k; idx++ {
			v := arr[idx]

			// 다음 첫 등장이 가장 나중인 값 찾기
			if plugged[v] && appearIndices[v] == -1 {
				appearIndices[v] = idx
			}
		}

		maxIdx := -1
		unplugTarget := -1
		for v := 1; v <= k; v++ {
			if !plugged[v] {
				continue
			}
			appearIdx := appearIndices[v]
			// 앞으로 한 번도 등장 안 하면 빼기
			if appearIdx == -1 {
				unplugTarget = v
				break
			}
			// 다음 첫 등장이 가장 나중인 값 빼기
			if appearIdx > maxIdx {
				maxIdx = appearIdx
				unplugTarget = v
			}
		}

		plugged[unplugTarget] = false
		plugged[curr] = true
		unplugCnt++
	}

	_, _ = fmt.Fprintf(writer, "%d", unplugCnt)
}
