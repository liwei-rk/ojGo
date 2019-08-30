## 孙悟空吃桃

```
孙悟空爱吃蟠桃，有一天趁着蟠桃园守卫不在来偷吃。
已知蟠桃园有N课桃树， 每棵树上都有桃子，天兵天将将在H小时后回来。
孙悟空可以决定他吃蟠桃的速度K（个/小时），每个小时选一棵桃树，从中吃掉K个，
如果树上的桃子少于K个，则全部吃掉，在这一小时剩余的时间里不再吃桃。
孙悟空喜欢慢慢吃，但又想在天兵天将回来前吃完所有桃子
求最小的K

输入
从标准输入读取一行数字， 前面数字表示每棵树上蟠桃个数，最后一个数字表示天兵天将离开的时间。

输出
吃掉所有桃子的最小速度K（整数），无解或输入异常时输出-1
```
朴素实现：O(n*max) n为桃树棵数，max为桃子对多那棵树上的桃子数
```
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
	for k := 1; k < max; k++ {
		if costHoursForK(trees, k) <= limitedHours {
			return k
		}
	}
	return -1
}

func costHoursForK(trees []int, k int) int {
	var costHours int
	for j := 0; j < len(trees); j++ {
		costHours += trees[j] / k
		if trees[j]%k > 0 {
			costHours ++
		}
	}
	return costHours
}
```
k从1到max去遍历，太老实了，如 2 3 800 5 这样的输入，有太多多余的尝试<br>
这里可以用二分查找的思路，将复杂度变成O(nlog(max))<br>
注意costHours <= limitedHours的情况及循环退出边界为left == k
```
func leastK(trees []int, max, limitedHours int) int {
	n := len(trees)
	if n > limitedHours {
		return -1
	}
	if n == limitedHours {
		return max
	}
	min := 1
	result := -1
	for {
		k := (min + max) / 2
		costHours := costHoursForK(trees, k)
		switch {
		case costHours == limitedHours:
			return k
		case costHours > limitedHours:
			min = k
		case costHours < limitedHours:
			max = k
			result = k
		}
		if min == k {
			break
		}
	}
	return result
}
```
还可以有一个优化，一开始的left其实可以是桃子总数/限定时间, 比这更小就无法在限定时间吃完了。
```
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
```
