package array

func linear_search(haystack []int, needle int) int {
  for i, element := range haystack {
    if element == needle {
      return i 
    }
  }
  return -1
}
