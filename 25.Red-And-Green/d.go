package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%v", &input)
	if len(input) <= 0 || len(input) >= 50 {
		panic("input is out of range")
	}
	fmt.Println(leastPaintBricks(input))
}

func leastPaintBricks(s string) int {
	const red = 'R'
	n := len(s)
	dp := make([]int, n)
	paintedReds := make([]int, n)
	if s[0] != red {
		paintedReds[0] = 1
	}

	for i := 1; i < n; i++ {
		if s[i] == red {
			paintedReds[i] = paintedReds[i-1]
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
			paintedReds[i] = paintedReds[i-1] + 1
		}
		// dp[i] = min(paintedReds[i], dp[i])
		if paintedReds[i] < dp[i] {
			dp[i] = paintedReds[i]
		}
	}
	return dp[n-1]
}
