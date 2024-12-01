package main

import "fmt"

func main() {
  list := []int{1, 2, 3, 4, 5}
  fmt.Println(binarySearch(list, 4))
}

// Binary search algorithm

func binarySearch(list []int, val int) int {
  low, high := 0, len(list) - 1

  for low <= high {
    mid := (low + high) / 2

    if list[mid] == val {
      return mid
    }

    if list[mid] > val {
      high = mid - 1
      continue
    }
    low = mid + 1
  }
  // not founded
  return -1
}
