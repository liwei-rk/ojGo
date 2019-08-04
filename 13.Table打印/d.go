package main

import "fmt"

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		return
	}
	if n <= 0 || n >= 13 || m <= 0 || m >= 13 {
		panic("iputed n, m is out of range")
		return
	}
	printTable(n, m)
}

func printTable(n, m int) {
	for i := 0; i < n; i++ {
		printPlusLine(m)
		printChanalLine(m)
		if i == n-1 {
			printPlusLine(m)
		}
	}
}

// 打印 +-- 行
func printPlusLine(m int) {
	for j := 0; j <= m; j++ {
		fmt.Print("+")
		if j < m {
			fmt.Print("---")
		}
	}
	fmt.Println()
}

// 打印竖线行
func printChanalLine(m int) {
	for j := 0; j <= m; j++ {
		fmt.Print("|")
		if j < m {
			fmt.Print("   ")
		}
	}
	fmt.Println()
}
