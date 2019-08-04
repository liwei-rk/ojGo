package main

import (
	"fmt"
	"math/big"
)

type dCase struct {
	a *big.Int
	b *big.Int
}

func main() {
	var cases = make([]dCase, 0)
	for {
		var a, b string
		if _, err := fmt.Scanf("%v %v", &a, &b); err != nil {
			break
		}
		if len(a) == 0 || len(b) == 0 {
			panic("invalid input, number cann't be empty")
		}
		if a[0] == '-' || b[0] == '-' {
			panic("invalid input, number must be positive")
		}
		aInt, _ := big.NewInt(0).SetString(a, 10)
		bInt, _ := big.NewInt(0).SetString(b, 10)
		cases = append(cases, dCase{a: aInt, b: bInt})
	}

	for _, c := range cases {
		fmt.Println(big.NewInt(0).Mul(c.a, c.b))
	}
}
