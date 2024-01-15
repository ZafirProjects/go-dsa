package array

import (
  "testing"
)

func TestCrystalBalls(t *testing.T)  {
  tests := []struct{
    floors []bool
    index int
  }{
    {[]bool{false, false, true, true}, 2},
    {[]bool{false, false, false, true}, 3},
    {[]bool{}, -1},
    {[]bool{true, true, true, true}, 0},
    {[]bool{false, false, false, false}, -1},
  }

  for i, test := range tests {
    if got := crystal_balls(test.floors); got != test.index {
      t.Fatalf("Failed test case #%d. Want %d got %d", i, test.index, got)
    }
  }
}
