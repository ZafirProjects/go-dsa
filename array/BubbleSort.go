package array

func bubble_sort(unsortedList []int) []int {
  for i := 0; i < len(unsortedList); i++ {
    for ii := 0; ii < len(unsortedList) - 1; ii++ {
      if unsortedList[ii] > unsortedList[ii + 1] {
        unsortedList[ii], unsortedList[ii + 1] = unsortedList[ii + 1], unsortedList[ii]
      }
    }
  }
  return unsortedList;
}
