package main

import (
	"fmt"
	"math"
)

func main() {
	var i uint
	for {
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			return
		}
		if IsConsecutive1(i) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
另一个实现，O(sqrt(2n))， 可打印序列

假设n可以表示成m个连续自然数的和
n = a + (a+1) +...+ (a+m-1) = (2*a+m-1)*m / 2 = am + m*(m-1)/2 (a > 0，否则就要从a+1开始）
可得a = [n - m*(m-1)/2] / m > 0
同样可得m的范围
*/
func IsConsecutive1(n uint) bool {
	for m := uint(2); m < uint(math.Sqrt(float64(n*2)))+1; m++ {
		s := n - m*(m-1)/2
		if s%m == 0 {
			a := s / m
			if a > 0 { // n == a + (a+1) +...+ i
				return true
			}
		}
	}
	return false
}
