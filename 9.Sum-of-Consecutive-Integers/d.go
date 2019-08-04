package main

import (
	"fmt"
)

func main() {
	var i uint
	for {
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			return
		}
		if IsConsecutive(i) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
从二进制角度去看，会很简单
一个正整数n，如果是2的幂(二进制中只有一个1)，则不满足题意，否则满足

因n < 2^30 + 1最多循环30次, 但无法打印序列
*/
func IsConsecutive(n uint) bool {
	oneNums := 0
	for n != 0 {
		if n&1 == 1 {
			oneNums++
		}
		n >>= 1
	}
	return oneNums > 1
}
