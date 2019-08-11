package main

import (
	"fmt"
	"math/big"
)

func main() {
	nums := []int{}
	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}
		if n <= 1 || n > 100 {
			break
		}
		nums = append(nums, n)
	}
	for _, n := range nums {
		fmt.Println(lcmN(n))
	}
}

func lcmN(n int) *big.Int {
	bigN := big.NewInt(int64(n))
	one := big.NewInt(1)
	r := big.NewInt(1)
	for bigN.Cmp(one) > 0 {
		lcm := lcm(bigN, r)
		if lcm.Cmp(r) > 0 {
			r = lcm
		}
		bigN.Sub(bigN, one)
	}
	return r
}

// lcm(a, b) == a * b / gcd(a, b)
// a and b must not be motified
func lcm(a, b *big.Int) *big.Int {
	t := big.NewInt(0).Set(a)
	gcd := t.GCD(nil, nil, a, b)
	t.Div(a, gcd)
	return t.Mul(t, b)
}

