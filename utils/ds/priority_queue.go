package ds

type Priority interface {
	Priority() int
}

func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{
		items: make([]Priority, 0, size),
	}
}

type PriorityQueue struct {
	items []Priority
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].Priority() < pq.items[j].Priority()
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(Priority)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}
