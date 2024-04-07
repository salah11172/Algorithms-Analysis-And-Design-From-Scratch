package main

import (
	"fmt"
)

type Item struct {
	Name   string
	Weight int
	Profit int
}

func main() {
	items := []Item{
		{Name: "#1", Weight: 1, Profit: 4},
		{Name: "#2", Weight: 3, Profit: 9},
		{Name: "#3", Weight: 5, Profit: 12},
		{Name: "#4", Weight: 4, Profit: 11},
	}

	maxWeight := 8

	dp := make([][]int, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		dp[i] = make([]int, maxWeight+1)
	}

	items = append([]Item{{Name: "#0", Weight: 0, Profit: 0}}, items...)

	for i := 1; i < len(items); i++ {
		for j := 0; j <= maxWeight; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if items[i].Weight <= j {
				dp[i][j] = max(dp[i-1][j], items[i].Profit+dp[i-1][j-items[i].Weight])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println("Max Profit:", dp[len(items)-1][maxWeight])

	solution := make([]string, 0)
	i := len(items) - 1
	j := maxWeight
	remain := maxWeight

	for remain >= 0 && j > 0 {
		if dp[i][j] > dp[i-1][j] {
			solution = append(solution, items[i].Name)
			remain = remain - items[i].Weight
			j = remain
			i--
		} else {
			i--
		}
	}

	fmt.Println(solution)
}
