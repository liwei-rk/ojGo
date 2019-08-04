## 连续自然数和
```
题目描述
2005年的百度之星初赛有这么一道题，一个正整数有可能可以被表示为 m(m>1) 个连续正整数之和，如：

15=1+2+3+4+5
15=4+5+6
15=7+8
但现在你的任务是判断给定的整数n能否表示成连续的m(m>1)个正整数之和。

解答要求
时间限制：1000ms, 内存限制：64MB
输入
输入只有一个整数n (1<n<2^30 +1)。

输出
若n能表示成连续的m(m>1)个正整数之和则输出“YES”，否则输出“NO”
```
1. 从二进制角度去看，会很简单<br>
一个正整数n，如果是2的幂(二进制中只有一个1)，则不满足题意，否则满足<br>
因n < 2^30 + 1最多循环30次, 但无法打印序列<br>
参考 d.go
```
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
```
2. 另一个实现，O(sqrt(2n))， 可打印序列<br>
假设n可以表示成m个连续自然数的和<br>
n = a + (a+1) +...+ (a+m-1) = (2*a+m-1)*m / 2 = am + m*(m-1)/2 (a > 0，否则就要从a+1开始）<br>
可得a = [n - m*(m-1)/2] / m > 0<br>
同样可得m的范围<br>
参考 e.go
```
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
```
