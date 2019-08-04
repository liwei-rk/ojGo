## 阶乘
```
题目描述
Given an integer N(0 ≤ N ≤ 10000), your task is to calculate N!.

解答要求
时间限制：5000ms, 内存限制：64MB
输入
One N in one line.

输出
For each N, output N! in one line.

样例
输入样例 1

3
输出样例 1

6
```
用循环和递归都可以，我们选循环：
```
func factorial(n uint) uint {
	result := uint(1)
	for i := uint(2); i <= n; i++ {
		result *= i
	}
	return result
}
```
不过数字一大就被截断了，还是要用big包：
```
func factorial(n int) *big.Int {
	r := big.NewInt(1)
	for i := 2; i <= n; i++ {
		r.Mul(r, big.NewInt(int64(i)))
	}
	return r
}
```
