package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	if n > 10000 || n < 0 {
		panic("n is out of range")
	}
	fmt.Println(factorial(n))
}

func factorial(n int) *big.Int {
	r := big.NewInt(1)
	for i := 2; i <= n; i++ {
		r.Mul(r, big.NewInt(int64(i)))
	}
	return r
}
