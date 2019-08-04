## 大数相乘

```
题目描述
I have a very simple problem for you. Given two integers A and B, your job is to multiply the product is of A × B.

解答要求
时间限制：1000ms, 内存限制：64MB
输入
Each line will contain two integers A and B. Process to end of file.
Notice that the integers are very large, that means you should not process them by using 32-bit integer.
You may assume the length of each integer will not exceed 1000.

输出
For each case, output the product is of A × B in one line.

样例
输入样例 1

1 2
3 11
14512451451245124512 15125125124512451245
输出样例 1

2
33
219502644063494817653152060344354417440
```
`注意这次的输入、输出和大数相加大数相减题目不一样`
```
1. 标准库big包的实现：d.go
2. 参考标准库，底层用数组实现
3. 模拟竖式计算，类似大数相加和相减的链表模拟实现
```
