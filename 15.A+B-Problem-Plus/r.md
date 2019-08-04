## 大数相加
```
题目描述
I have a very simple problem for you. Given two integers A and B, your job is to calculate the Sum of A + B.

解答要求
时间限制：1000ms, 内存限制：64MB
输入
The first line of the input contains an integer T(1≤T≤20) which means the number of test cases.
Then T lines follow, each line consists of two positive integers, A and B. Notice that the integers are very large,that means you should not process them by using 32-bit integer.You may assume the length of each integer will not exceed 1000.

输出
For each test case, you should output two lines. The first line is "Case #:", # means the number of the test case. The second line is the an equation "A + B = Sum", Sum means the result of A + B.Note there are some spaces int the equation. Output a blank line between two test cases.

样例
输入样例 1 复制

2
1 2
112233445566778899 998877665544332211
输出样例 1

Case 1:
1 + 2 = 3

Case 2:
112233445566778899 + 998877665544332211 = 1111111111111111110
```
```
1. 使用标准库big包: d.go
2. 模拟big包实现，底层用一个uint32的数组
3. 就加法而言，使用链表模拟大数，将每一位倒序存入，注意进位: e.go
```
