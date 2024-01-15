package array

func qs(array []int, lo, hi int) []int {
  if lo >= hi {
    return nil
  }

  pivotIdx := partition(array, lo, hi)

  qs(array, lo, pivotIdx - 1)
  qs(array, pivotIdx + 1, hi)

  return array
}

func partition(array []int, lo, hi int) int {
  pivot := array[hi]
  idx := lo
  for i := lo; i < hi; i++ {
    if array[i] <= pivot {
      array[i], array[idx] = array[idx], array[i]
      idx++
    }
  }
  array[hi], array[idx] = array[idx], array[hi]
  return idx
}

func quick_sort(unsortedList []int) []int {
  sorted := qs(unsortedList, 0, len(unsortedList) - 1)
  return sorted
}

