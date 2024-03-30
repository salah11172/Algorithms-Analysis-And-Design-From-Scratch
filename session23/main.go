package main

import (
	"fmt"
	"math"
)

func main() {
	labels := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	data := [][]int{
		{0, 2, 4, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 4, 6, 0, 0, 0},
		{0, 0, 0, 0, 3, 2, 4, 0, 0, 0},
		{0, 0, 0, 0, 4, 1, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 4, 0},
		{0, 0, 0, 0, 0, 0, 0, 6, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 3, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	n := len(data)
	states := make([]struct {
		from string
		to   string
		cost int
	}, n)

	for i := n - 2; i >= 0; i-- {
		states[i].from = labels[i]
		states[i].cost = math.MaxInt64
		for j := i + 1; j < n; j++ {
			if data[i][j] == 0 {
				continue
			}

			newCost := data[i][j] + states[j].cost

			if newCost < states[i].cost {
				states[i].to = labels[j]
				states[i].cost = newCost
			}
		}
	}

	path := []string{"A"}
	i, j := 0, 0
	for i < len(states) {
		if states[i].from == path[j] {
			path = append(path, states[i].to)
			j++
		}
		i++
	}

	fmt.Println("Minimum Cost:", states[0].cost)
	fmt.Println("Minimum Path:", path)
}
