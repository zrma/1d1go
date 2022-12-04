package ds

type UnionFind struct {
	// parents 배열은 자신이 루트일 경우 트리의 크기를 음수로 저장
	//         자식일 경우 부모의 인덱스 저장
	parents []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n+1)
	for i := 0; i <= n; i++ {
		// 각각 자신을 루트로 설정, 트리의 크기 == 1개 (== -1)
		parents[i] = -1
	}
	return &UnionFind{parents: parents}
}

func (uf *UnionFind) Find(a int) int {
	// 자신이 루트일 경우(트리의 크기: 음수) 자신을 반환
	if uf.parents[a] < 0 {
		return a
	}
	uf.parents[a] = uf.Find(uf.parents[a])
	return uf.parents[a]
}

func (uf *UnionFind) Size(a int) int {
	return -uf.parents[uf.Find(a)]
}

func (uf *UnionFind) Children() []int {
	children := make([]int, 0, len(uf.parents))
	for _, p := range uf.parents[1:] {
		if p < 0 {
			children = append(children, -p)
		}
	}
	return children
}

func (uf *UnionFind) Union(a int, b int) {
	a = uf.Find(a)
	b = uf.Find(b)

	if a == b {
		return
	}

	// 각 트리의 루트는 음수
	// 값이 작은 경우가 트리의 크기가 더 크다.
	// 트리의 크기가 더 큰 곳에 붙인다. (트리가 더 깊어지지 않도록)
	if uf.parents[a] <= uf.parents[b] {
		uf.parents[a] += uf.parents[b]
		uf.parents[b] = a
	} else {
		uf.parents[b] += uf.parents[a]
		uf.parents[a] = b
	}
}
