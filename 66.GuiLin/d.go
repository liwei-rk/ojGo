package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		panic(err)
	}
	if n < 1 || n > 100 {
		panic("input out of range")
	}
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n+1)
		for j := i + 1; j <= n; j++ {
			fmt.Scan(&cost[i][j])
		}
	}
	fmt.Println(minCost(cost))
}

func minCost(cost [][]int) int {
	n := len(cost)
	costMin := make([]int, n+1)
	costMin[1] = cost[0][1]
	for i := 1; i <= n; i++ {
		costMin[i] = cost[0][i] // 直接从第0站到第i站的租金
		for k := 1; k < i; k++ {
			kCost := costMin[k] + cost[k][i] // 从第0站到第k站的最小租金 + 从第k站直接到第i站的租金
			if costMin[i] > kCost {
				costMin[i] = kCost
			}
		}
	}
	return costMin[n]
}
