package array

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T)  {
  tests := []struct{
    unsortedList []int
    sortedList []int
  }{
    {[]int{1,2,4,3,5}, []int{1,2,3,4,5}},
    {[]int{1,2,9,5,7}, []int{1,2,5,7,9}},
    {[]int{}, []int{}},
    {[]int{1}, []int{1}},
  } 
  for i, test := range tests {
    got := bubble_sort(test.unsortedList)
    if !slices.Equal(got, test.sortedList) {
      t.Fatalf("Failed test case #%d. Want %d got %d", i, test.sortedList, got)
    }
  }
}
