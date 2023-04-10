package queueu

import "fmt"

type Queue[T any] struct {
	nodes []T
}

func (q *Queue[T]) Add(val T) {
	q.nodes = append(q.nodes, val)
}

func (q *Queue[T]) Remove() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, fmt.Errorf("queue is empty")
	}

	result := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]

	return result, nil
}

func (q *Queue[T]) Size() int {
	return len(q.nodes)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
