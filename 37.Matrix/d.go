package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	if n < 1 || n > 100 || m < 1 || m > 100 {
		panic("input out of range")
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	fmt.Println(MaxSum(matrix, n, m))
}

func MaxSum(matrix [][]int, n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = matrix[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + matrix[i][j]
				}
			}
		}
	}
	return dp[n-1][m-1]
}
