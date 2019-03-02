package tutorial_30_days_of_code

type queue struct {
	nodes []string
}

func (q *queue) enqueue(item string) {
	q.nodes = append(q.nodes, item)
}

func (q *queue) dequeue() (string, bool) {
	if len(q.nodes) < 1 {
		return "", false
	}
	item := q.nodes[0]
	q.nodes = q.nodes[1:]

	return item, true
}

type stack struct {
	nodes []string
}

func (s *stack) push(item string) {
	s.nodes = append(s.nodes, item)
}

func (s *stack) pop() (string, bool) {
	if len(s.nodes) < 1 {
		return "", false
	}
	item := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]

	return item, true
}
