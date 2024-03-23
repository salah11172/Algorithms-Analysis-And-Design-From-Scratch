package main

import "fmt"

func mergeSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	}

	middle := length / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	arr := []int{6, 16, 13, 5, 1, 7, 20, 3}
	fmt.Println("Unsorted array:", arr)
	sorted := mergeSort(arr)
	fmt.Println("Sorted array:", sorted)
}
