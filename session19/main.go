package main

import (
	"fmt"
)

type CharFreq struct{}

func (cf *CharFreq) ASCIIMethod(message string) {
	fmt.Println("ASCIIMethod")

	charFrequencyMap := make(map[int]int)

	// Calculate frequencies using ASCII codes
	for _, char := range message {
		currentCode := int(char)
		charFrequencyMap[currentCode]++
	}

	for code, frequency := range charFrequencyMap {
		fmt.Printf("%c %d\n", code, frequency)
	}

	cf.SortHash(charFrequencyMap)

}

func (cf *CharFreq) SortHash(charFrequencyMap map[int]int) {
	freqArray := make([][2]int, 0, len(charFrequencyMap))

	// Convert the character frequencies map to an array
	for char, frequency := range charFrequencyMap {
		freqArray = append(freqArray, [2]int{int(char), frequency})
	}

	cf.sort(freqArray, 0, len(freqArray)-1)

	fmt.Println("Print Sorted data ...")
	for _, entry := range freqArray {
		fmt.Printf("%c %d\n", rune(entry[0]), entry[1])
	}
}

func (cf *CharFreq) sort(array [][2]int, start, end int) {
	if end <= start {
		return
	}

	midpoint := (end + start) / 2
	cf.sort(array, start, midpoint)
	cf.sort(array, midpoint+1, end)
	cf.merge(array, start, midpoint, end)
}

func (cf *CharFreq) merge(array [][2]int, start, mid, end int) {
	leftLength := mid - start + 1
	rightLength := end - mid

	leftArray := make([][2]int, leftLength)
	rightArray := make([][2]int, rightLength)

	for i := 0; i < leftLength; i++ {
		leftArray[i] = array[start+i]
	}
	for j := 0; j < rightLength; j++ {
		rightArray[j] = array[mid+1+j]
	}

	i, j, k := 0, 0, start
	for i < leftLength && j < rightLength {
		if leftArray[i][1] <= rightArray[j][1] {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
		k++
	}
	for i < leftLength {
		array[k] = leftArray[i]
		i++
		k++
	}
	for j < rightLength {
		array[k] = rightArray[j]
		j++
		k++
	}
}

func main() {
	charFreq := CharFreq{}

	message := "hello   world"
	charFreq.ASCIIMethod(message)
}
