package array

import(
  "testing"
  "slices"
)

func TestQuickSort(t *testing.T)  {
  tests := []struct{
    unsorted, sorted []int
  }{
    {[]int{9,3,7,4,69,420,42}, []int{3,4,7,9,42,69,420}},
  }

  for i, test := range tests {
    if got := quick_sort(test.unsorted); !slices.Equal(got, test.sorted){
      t.Fatalf("Failed test %d, exptected %d but got %d", i, test.sorted, got)
    }
  }
}
