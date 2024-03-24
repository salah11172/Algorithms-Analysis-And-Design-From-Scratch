package main

import (
	"fmt"
)

// binarySearch searches for target in the sorted array arr
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Check if target is present at mid
		if arr[mid] == target {
			return mid
		}

		// If target is greater, ignore left half
		if arr[mid] < target {
			left = mid + 1
		} else { // If target is smaller, ignore right half
			right = mid - 1
		}
	}

	// If we reach here, then the element was not present
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	target := 10
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
