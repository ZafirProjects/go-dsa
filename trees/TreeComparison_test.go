package trees

import ("testing")

func TestCompareBinaryTree(t *testing.T)  {
  bt1 := binary_tree_constructor(50)
  bt1.addItem(20)
  bt1.addItem(30)
  bt1.addItem(70)
  bt1.addItem(100)

  bt2 := binary_tree_constructor(50)
  bt2.addItem(20)
  bt2.addItem(30)
  bt2.addItem(70)
  bt2.addItem(100)

  if got := compare_binary_tree(&bt1, &bt2); !got {
    t.Fatalf("failed test, expected true, got false")
  }
}
