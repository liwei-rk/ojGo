package main

import (
	"fmt"
)

func main() {
	var input []int
	var max int
	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}
		input = append(input, n)
		if max < n {
			max = n
		}
	}
	n := len(input) - 1
	limitedHours := input[n]
	trees := input[:n]
	fmt.Println(leastK(trees, max, limitedHours))
}

func leastK(trees []int, max, limitedHours int) int {
	n := len(trees)
	if n > limitedHours {
		return -1
	}
	if n == limitedHours {
		return max
	}
	total := 0
	for _, peaches := range trees {
		total += peaches
	}
	min := total / limitedHours
	if min == 0 {
		min = 1
	}
	result := -1
	for {
		fmt.Printf("[%d, %d]\n", min, max)
		k := (min + max) / 2
		costHours := costHoursForK(trees, k)
		switch {
		case costHours <= limitedHours:
			max = k
			result = k
		case costHours > limitedHours:
			min = k
		}
		if min == k {
			break
		}
	}
	return result
}

func costHoursForK(trees []int, k int) int {
	var costHours int
	for j := 0; j < len(trees); j++ {
		costHours += trees[j] / k
		if trees[j]%k > 0 {
			costHours ++
		}
	}
	fmt.Println("k:", k, "cost:", costHours)
	return costHours
}
