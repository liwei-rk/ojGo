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
	var t int
	if _, err := fmt.Scanf("%d", &t); err != nil {
		return
	}
	if t < 1 || t > 20 {
		return
	}
	var cases = make([]dCase, t)
	for i := 0; i < t; i++ {
		var a, b string
		if _, err := fmt.Scanf("%v %v", &a, &b); err != nil {
			panic("invalid input")
		}
		if len(a) == 0 || len(b) == 0 {
			panic("invalid input, number cann't be empty")
		}
		if a[0] == '-' || b[0] == '-' {
			panic("invalid input, number must be positive")
		}
		cases[i].a, _ = big.NewInt(0).SetString(a, 10)
		cases[i].b, _ = big.NewInt(0).SetString(b, 10)
	}

	for i, c := range cases {
		fmt.Printf("Case %d:\n", i+1)
		fmt.Printf("%v + %v = %v\n\n", c.a, c.b, big.NewInt(0).Add(c.a, c.b))
	}
}
