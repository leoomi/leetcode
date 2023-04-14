package stack

type Stack[T any] interface {
	Add(val T)
	Pop() T
	Size() int
}

type node[T any] struct {
	val  T
	next *node[T]
}

type linkedListStack[T any] struct {
	size int
	root *node[T]
}

func NewLinkedListStack[T any]() Stack[T] {
	return &linkedListStack[T]{}
}

func NewArrayStack[T any](size int) Stack[T] {
	s := &arrayStack[T]{}
	s.array = make([]T, size)
	s.capacity = size
	s.top = -1

	return s
}

func (s *linkedListStack[T]) Add(val T) {
	n := node[T]{
		val:  val,
		next: s.root,
	}

	s.root = &n
	s.size++
}

func (s *linkedListStack[T]) Pop() T {
	if s.root == nil {
		panic("stack is empty")
	}

	val := s.root.val
	s.root = s.root.next

	s.size--

	return val
}

func (s *linkedListStack[T]) Size() int {
	return s.size
}

type arrayStack[T any] struct {
	top      int
	capacity int
	array    []T
}

func (s *arrayStack[T]) Add(val T) {
	if s.top == s.capacity-1 {
		panic("stack is full")
	}

	s.top++
	s.array[s.top] = val
}

func (s *arrayStack[T]) Pop() T {
	if s.top == -1 {
		panic("stack is empty")
	}

	val := s.array[s.top]
	s.top--

	return val
}

func (s *arrayStack[T]) Size() int {
	return s.top + 1
}
