package p17400

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve17478(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	//goland:noinspection SpellCheckingInspection
	const msg = "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다."
	_, _ = fmt.Fprintln(writer, msg)

	recursive(writer, n, 0)
}

func recursive(writer Writer, n, cur int) {
	const underlines = "____"
	ss := make([]string, cur)
	for i := range ss {
		ss[i] = underlines
	}
	prefix := strings.Join(ss, "")

	_, _ = fmt.Fprintf(writer, "%s\"%s\"\n", prefix, "재귀함수가 뭔가요?")
	defer func() { _, _ = fmt.Fprintf(writer, "%s%s\n", prefix, "라고 답변하였지.") }()

	if cur == n {
		_, _ = fmt.Fprintf(writer, "%s\"%s\"\n", prefix, "재귀함수는 자기 자신을 호출하는 함수라네")
		return
	}

	_, _ = fmt.Fprintf(writer, "%s\"%s\n", prefix, "잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.")
	_, _ = fmt.Fprintf(writer, "%s%s\n", prefix, "마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.")
	_, _ = fmt.Fprintf(writer, "%s%s\"\n", prefix, "그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.")

	recursive(writer, n, cur+1)
}
