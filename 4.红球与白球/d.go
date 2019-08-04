package main

import "fmt"

func main() {
	var m, n, x int
	for {
		if _, err := fmt.Scanf("%d %d %d", &m, &n, &x); err != nil {
			return
		}
		if m == 0 && n == 0 && x == 0 {
			return
		}
		fmt.Printf("%.4f\n", p(m, n, x))
	}
}

// 红球多于白球 <=> 红球个数最少要x/2+1
func p(m, n, x int) float64 {
	return pro(m, n, x, x/2+1)
}

/*
 在装有m个红球，n个白球的袋中取出x个球，至少含有leastM个红球的概率
*/
func pro(m, n, x, leastM int) float64 {
	// 递归退出条件
	switch {
	case leastM > x || leastM > m:
		return 0
	case leastM == 0 || x-n > leastM:
		return 1
	}

	// 分两种情况，第一个取出的是红球/白球
	// 第一次取出红球的概率：m/(m+n)，接下来从m-1个红球和n个白球里取出x-1个球，其中红球至少leastM-1个，递归
	p1 := float64(m) / float64(m+n) * pro(m-1, n, x-1, leastM-1)
	// 第一次取出白球的概率：n/(m+n)，接下来从m个红球和n-1个白球里取出x-1个球，其中红球至少leastM个，递归
	p2 := float64(n) / float64(m+n) * pro(m, n-1, x-1, leastM)

	return p1 + p2
}
