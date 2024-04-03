package main

import (
	"fmt"
)

func main() {
	text01 := "HELLOWORLD"
	text02 := "OHELOD"
	n := len(text01)
	m := len(text02)

	text01 = " " + text01
	text02 = " " + text02

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if text02[i] == text01[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	// Constructing the LCS
	str := ""
	i := m
	j := n
	for i > 0 && j > 0 {
		if dp[i][j] > dp[i][j-1] {
			if dp[i][j] == dp[i-1][j] {
				i--
			} else {
				str = string(text02[i]) + str
				i--
				j--
			}
		} else {
			j--
		}
	}

	fmt.Println(dp[m][n])
	fmt.Println(str)
}
