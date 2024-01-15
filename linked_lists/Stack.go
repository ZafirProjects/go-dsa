package linkedlists

type SNode[T any] struct {
	value T
	prev  *SNode[T]
}
type Stack[T any] struct {
	length int
	head   *SNode[T]
}

func s_constructor() Stack[any] {
	s := Stack[any]{0, nil}
	return s
}

func (s *Stack[T]) push(item T) *Stack[T] {
	n := &SNode[T]{item, nil}
	if s.head == nil {
		s.head = n
		s.length += 1
		return s
	}

	s.length += 1
	n.prev = s.head
	s.head = n
	return s
}

func (s *Stack[T]) pop() *T {
	if s.head == nil {
		s.length -= 1
		return nil
	}

	value := s.head.value
	s.head = s.head.prev
	s.length -= 1

	return &value
}

func (s *Stack[T]) peek() T {
	return s.head.value
}
