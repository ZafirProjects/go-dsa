package linkedlists

import (
  "testing"
)

func TestQueue(t *testing.T)  {
  queue := q_constructor()
  queue.enqueue("item1")
  if queue.head.value != "item1" && queue.tail.value != "item1" && queue.length != 1 {
    t.Fatalf("Failed, wanted 'item1' but got %d and %d", queue.head.value, queue.tail.value)
  }
  queue.enqueue("item2")
  if queue.head.value != "item1" && queue.tail.value != "item2" && queue.length != 2 {
    t.Fatalf("Failed wanted head to be 'item1' and tail to be item2 but got %d and %d and length %d", queue.head.value, queue.tail.value, queue.length)
  }
  queue.dequeue()
  if queue.head.value != "item2" && queue.tail.value != "item2" && queue.length != 1 {
    t.Fatalf("Failed wanted head to be 'item2' and tail to be item2 but got %d and %d and length %d", queue.head.value, queue.tail.value, queue.length)
  }
  if got := queue.peek(); got != "item2" {
    t.Fatalf("Failed peek, expected item2, got %d", got)
  }
}
