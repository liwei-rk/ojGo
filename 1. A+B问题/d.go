package main

import "fmt"

func main() {
	var a, b int
	for {
		if _, err := fmt.Scanf("%d %d", &a, &b); err != nil {
			return
		} else {
			fmt.Println(a + b)
		}
	}
}
