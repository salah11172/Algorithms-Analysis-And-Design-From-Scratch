package main

import "fmt"

func insertionSort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Unsorted array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
