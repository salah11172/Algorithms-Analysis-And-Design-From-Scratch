package main

import (
	"fmt"
	"math"
)

func main() {
	labels := []rune{'1', '2', '3', '4', '5', '6'}
	graph := [][]float64{
		{0, 6.7, 5.2, 2.8, 5.6, 3.6},
		{6.7, 0, 5.7, 7.3, 5.1, 3.2},
		{5.2, 5.7, 0, 3.4, 8.5, 4.0},
		{2.8, 7.3, 3.4, 0, 8, 4.4},
		{5.6, 5.1, 8.5, 8, 0, 4.6},
		{3.6, 3.2, 4, 4.4, 4.6, 0},
	}

	v := 6

	selectedEdgesCount := 0
	selected := make([]bool, v)
	selected[0] = true

	for selectedEdgesCount < v-1 {
		minValue := math.MaxFloat64
		var tempFrom, tempTo int
		for i := 0; i < v; i++ {
			if selected[i] {
				for j := 0; j < v; j++ {
					if !selected[j] && graph[i][j] > 0 && graph[i][j] < minValue {
						minValue = graph[i][j]
						tempFrom = i
						tempTo = j
					}
				}
			}
		}

		selected[tempTo] = true
		selectedEdgesCount++

		fmt.Printf("%c - %c : %.1f\n", labels[tempFrom], labels[tempTo], graph[tempFrom][tempTo])
	}
}
