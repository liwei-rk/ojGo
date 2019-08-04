##  第一步
```
题目描述
Calculate A + B，and give me the answer!

解答要求
时间限制：1000ms, 内存限制：64MB
输入
Input two integers A and B, process to the end of the file. (Watch the Sample Input)

输出
For each case, output A + B in one line.(Watch the Sample Output)

输入样例：
2 3
3 4

输出样例：
2 3
5
3 4
7

提示
Note:

Press the Submit to submit your code!
Q: Where is the input and the output?
A: Your program shall read input from stdin('Standard Input') and write output to stdout('Standard Output').

For example, you can use 'scanf' in C or 'cin' in C++ to read from stdin, and use 'printf' in C or 'cout' in C++ to write to stdout.
User programs are not allowed to open and read from/write to files, you will get a "Restricted Function" if you try to do so .

Here is a sample solution for problem 1 using C:

#include <stdio.h>
int main(void)
{
    int a, b;
    while (scanf("%d%d", &a, &b) != EOF) {
        printf("%d\n", a + b);
    }
    return 0;
}
```
```
这里和大家做了输入输出的约定，注意要支持循环多组用例输入
```
