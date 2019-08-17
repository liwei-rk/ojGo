package main

import "fmt"

func main() {
	var inputs []int
	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}
		inputs = append(inputs, n)
	}
	for _, n := range inputs {
		fmt.Println(f(n))
	}
}

func f(n int) int {
	if n <= 0 || n > 50 {
		panic("input out of range")
	}
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-3]
	}
	return dp[n-1]
}
