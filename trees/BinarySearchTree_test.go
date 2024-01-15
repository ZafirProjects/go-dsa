package trees

import(
  "testing"
)

func TestBinaryTree(t *testing.T)  {
  bt := binary_tree_constructor(3)
  bt.addItem(1)
  if bt.left.value != 1 {
    t.Fatalf("Left Node should've been 1")
  }
  bt.addItem(2)
  if bt.left.right.value != 2 {
    t.Fatalf("Left Right Node should've been 2")
  }
  bt.addItem(4)
  if bt.right.value != 4 {
    t.Fatalf("Right Node should've been 4")
  }
  bt.addItem(5)
  if bt.right.right.value != 5 {
    t.Fatalf("Right Right Node should've been 5")
  }
  
}
