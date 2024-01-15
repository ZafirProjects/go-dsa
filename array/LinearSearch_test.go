package array

import (
  "testing"
)

func TestLinearSearch(t *testing.T)  {
  tests := []struct {
    haystack []int
    needle, index int
  }{
    {[]int{}, 3, -1},
    {[]int{1,2,3}, 3, 2},
    {[]int{1,2,8,9,78,12,45}, 3, -1},
    {[]int{1,2,8,9,78,12,45}, 12, 5},
    {[]int{1,2,9}, 9, 2},
  }

  for i, test := range tests {
    if got := linear_search(test.haystack, test.needle); got != test.index {
      t.Fatalf("Failed test case #%d. Want %d got %d", i, test.index, got)
    }
  }
}
