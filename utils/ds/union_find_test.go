package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnionFind(t *testing.T) {
	const size = 123
	uf := NewUnionFind(size)
	assert.Len(t, uf.parents, size+1)

	for i := 0; i <= size; i++ {
		assert.Equal(t, -1, uf.parents[i])
	}

	for i := 0; i <= size; i++ {
		assert.Equal(t, i, uf.Find(i))
	}
}

func TestFind(t *testing.T) {
	for _, tt := range []struct {
		give        int
		wantVal     int
		wantParents []int
	}{
		{1, 1, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
		{2, 1, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
		{3, 1, []int{-1, -1, 1, 1, 3, -1, 5, 6, 7, -1, -1}},
		{4, 1, []int{-1, -1, 1, 1, 1, -1, 5, 6, 7, -1, -1}},
		{5, 5, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
		{6, 5, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
		{7, 5, []int{-1, -1, 1, 2, 3, -1, 5, 5, 7, -1, -1}},
		{8, 5, []int{-1, -1, 1, 2, 3, -1, 5, 5, 5, -1, -1}},
		{9, 9, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
		{10, 10, []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}},
	} {
		t.Run(fmt.Sprintf("%d", tt.give), func(t *testing.T) {
			give := []int{-1, -1, 1, 2, 3, -1, 5, 6, 7, -1, -1}
			uf := &UnionFind{parents: give}
			assert.Equal(t, tt.wantVal, uf.Find(tt.give))
			assert.Equal(t, tt.wantParents, uf.parents)
		})
	}
}

func TestUnionFind_Size(t *testing.T) {
	uf := &UnionFind{parents: []int{0, -5, 1, 2, 3, 4, 5}}
	assert.Equal(t, 5, uf.Size(6))
	assert.Equal(t, []int{0, -5, 1, 1, 1, 1, 1}, uf.parents)
}

func TestUnion(t *testing.T) {
	uf := NewUnionFind(10)
	assert.Equal(t, []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, uf.parents)

	uf.Union(1, 2)
	assert.Equal(t, []int{-1, -2, 1, -1, -1, -1, -1, -1, -1, -1, -1}, uf.parents)
	uf.Union(2, 3)
	assert.Equal(t, []int{-1, -3, 1, 1, -1, -1, -1, -1, -1, -1, -1}, uf.parents)
	uf.Union(3, 4)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -1, -1, -1, -1, -1, -1}, uf.parents)

	uf.Union(5, 6)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -2, 5, -1, -1, -1, -1}, uf.parents)
	uf.Union(6, 7)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -3, 5, 5, -1, -1, -1}, uf.parents)

	uf.Union(8, 9)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -3, 5, 5, -2, 8, -1}, uf.parents)
	uf.Union(9, 10)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -3, 5, 5, -3, 8, 8}, uf.parents)

	uf.Union(6, 10)
	assert.Equal(t, []int{-1, -4, 1, 1, 1, -6, 5, 5, 5, 8, 8}, uf.parents)

	uf.Union(3, 8)
	assert.Equal(t, []int{-1, 5, 1, 1, 1, -10, 5, 5, 5, 8, 8}, uf.parents)
}

func TestUnion_EqualVal(t *testing.T) {
	want := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	uf := NewUnionFind(10)

	for i := 0; i <= 10; i++ {
		uf.Union(i, i)
		assert.Equal(t, want, uf.parents)
	}
}
