package array

import (
  "math"
)

func crystal_balls(floors []bool) int {
  numberOfFloors := len(floors)

  if numberOfFloors == 0 || floors[numberOfFloors - 1] == false {
    return -1
  }

  pointer := 0
  jumpAmount := math.Sqrt(float64(numberOfFloors))

  for ; pointer < numberOfFloors; pointer += int(jumpAmount) {
    if floors[pointer] {
      return pointer
    }
  }

  pointer -= int(jumpAmount)
  for i := pointer; i < numberOfFloors; i++ {
    if floors[i] {
      return i
    }
  }

  return -1
}
