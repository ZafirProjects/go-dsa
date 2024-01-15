package trees

import(
  "testing"
  "slices"
)

func TestPreOrderSearch(t *testing.T)  {
  bt := binary_tree_constructor(50)
  bt.addItem(25)
  bt.addItem(75)
  bt.addItem(10)
  bt.addItem(30)
  bt.addItem(60)
  bt.addItem(100)
  solution := []int{50, 25, 10, 30, 75, 60, 100} 
  if got := pre_order_search(&bt); !slices.Equal(solution, got) {
    t.Fatalf("Test failed, expected %d, got %d", solution, got)
  }
}
