package main

import (
	"fmt"
)

// Merge function
func merge(arr []int, leftStart, leftEnd, rightStart, rightEnd int) {
	left := make([]int, leftEnd-leftStart+1)
	right := make([]int, rightEnd-rightStart+1)

	copy(left, arr[leftStart:leftEnd+1])
	copy(right, arr[rightStart:rightEnd+1])

	i, j := 0, 0
	k := leftStart

	for i < len(left) && left[i] < 0 {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) && right[j] < 0 {
		arr[k] = right[j]
		j++
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// MergeSort function
func mergeSort(arr []int, start, end int) {
	if start < end {
		mid := start + (end-start)/2
		mergeSort(arr, start, mid)
		mergeSort(arr, mid+1, end)
		merge(arr, start, mid, mid+1, end)
	}
}

// Segregate function
func segregate(arr []int) []int {
	mergeSort(arr, 0, len(arr)-1)
	return arr
}

func main() {
	arr := []int{3, -5, 1, -2, 7, -8, 9, -4}
	fmt.Println("Original array:", arr)

	arr = segregate(arr)
	fmt.Println("Segregated array:", arr)
}
