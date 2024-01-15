package trees

import("testing")

func TestBFS(t *testing.T)  {
  bt := binary_tree_constructor(50)
  bt.addItem(1)
  bt.addItem(10)
  bt.addItem(25)
  bt.addItem(30)
  bt.addItem(45)
  bt.addItem(75)
  bt.addItem(100)
  tests := []struct{
    needle int
    solution bool
  }{
    {1, true},
    {5, false},
    {25, true},
    {30, true},
    {70, false},
    {100, true},
  }

  for i, test := range tests {
    if got := breadth_first_search(bt, test.needle); got != test.solution {
      t.Fatalf("failed test %d. expected %t, got %t", i, test.solution, got)
    }
  }
}
