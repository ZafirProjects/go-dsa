package trees

import("testing")

func TestDfsOnBst(t *testing.T)  {
  bt := binary_tree_constructor(100)
  bt.addItem(10)
  bt.addItem(17)
  bt.addItem(30)
  bt.addItem(50)
  bt.addItem(64)
  bt.addItem(72)
  bt.addItem(99)
  bt.addItem(101)
  bt.addItem(121)
  bt.addItem(199)

  tests := []struct{
    needle int
    solution bool
  }{
    {10, true},
    {15, false},
    {50, true},
    {72, true},
    {90, false},
    {100, true},
    {199, true},
    {200, false},
  }

  for i, test := range tests {
    if got := dfs_bst(bt, test.needle); got != test.solution {
      t.Fatalf("Failed test %d, expected %t but got %t", i, test.solution, got)
    }
  }
}
