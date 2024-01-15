package linkedlists

import (
  "testing"
)

func TestStack(t *testing.T)  {
  stack := s_constructor()
  stack.push("item1")
  if stack.head.value != "item1" && stack.length != 1 {
    t.Fatalf("Failed push, wanted 'item1' but got %d", stack.head.value)
  }
  stack.push("item2")
  if stack.head.value != "item2" && stack.length != 2 {
    t.Fatalf("Failed push, wanted 'item2' but got %d", stack.head.value)
  }
  stack.pop()
  if stack.head.value != "item1" && stack.length != 1 {
    t.Fatalf("Failed pop, wanted 'item1' but got %d", stack.head.value)
  }
  if got := stack.peek(); got != "item1" {
    t.Fatalf("Failed peek, wanted 'item1' but got %d", got)
  }
}
