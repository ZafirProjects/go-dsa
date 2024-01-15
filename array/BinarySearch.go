package array

func binary_search(haystack []int, needle int) int {
  lo := 0
  hi := len(haystack)

  for lo < hi {
    mid := lo + (hi - lo) / 2
    value := haystack[mid]
    if value == needle {
      return mid
    }else if value > needle {
      hi = mid
    }else {
      lo = mid + 1
    }
  }
  return -1
}
