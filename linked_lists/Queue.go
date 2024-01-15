package linkedlists

type QNode[T any] struct{
  value T
  next *QNode[T]
}

type Queue[T any] struct{
  length int
  head, tail *QNode[T]
}

func q_constructor() Queue[any] {
  q := Queue[any]{0, nil, nil}
  return q
}

func (q *Queue[T]) enqueue(item T) *Queue[T] {
  n := &QNode[T]{item, nil}
  if q.head == nil {
    q.tail = n
    q.head = n
    return q
  }

  q.tail.next = n
  q.tail = n
  q.length += 1
  return q
}

func (q *Queue[T]) dequeue() *T {
  if q.head == nil {
    return nil
  }

  q.length -= 1

  head := q.head
  q.head = q.head.next
  return &head.value
}

func (q *Queue[T]) peek() T {
  return q.head.value
}

