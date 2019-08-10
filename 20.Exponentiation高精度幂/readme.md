## 大数高精度幂次
```
题目描述
Problems involving the computation of exact values of very large magnitude and precision are common. 
For example, the computation of the national debt is a taxing experience for many computer systems.
This problem requires that you write a program to compute the exact value of R^n 
where R is a real number ( 0.0 < R < 99.999 ) and n is an integer such that 0 < n ≤ 25.

解答要求
时间限制：4999ms, 内存限制：64MB
输入
The input will consist of a set of pairs of values for R and n.
The R value will occupy columns 1 through 6, and the n value will be in columns 8 and 9.

输出
The output will consist of one line for each line of input giving the exact value of Rn.
Leading zeros should be suppressed in the output.
Insignificant trailing zeros must not be printed. Don't print the decimal point if the result is an integer.

样例
输入样例 1

95.123 12
0.4321 20
5.1234 15
6.7592  9
98.999 10
1.0100 12
输出样例 1

548815620517731830194541.899025343415715973535967221869852721
.00000005148554641076956121994511276767154838481760200726351203835429763013462401
43992025569.928573701266488041146654993318703707511666295476720493953024
29448126.764121021618164430206909037173276672
90429072743629540498.107596019456651774561044010001
1.126825030131969720661201
```
标准库big包没有power函数~~<br>
先实现以下power函数
```
func power(r string, n int) string {
	bigR, _ := big.NewFloat(0).SetString(r)
	bigResult := big.NewFloat(0)
	bigResult.Add(bigResult, bigR)
	for i := 1; i < n; i++ {
		bigResult.Mul(bigResult, bigR)
	}
	result := bigResult.Text('f', 100)
	if strings.Contains(result,".") {
		result = strings.TrimRight(result, "0")
		if strings.HasSuffix(result, ".") {
			result = result[:len(result)-1]
		}
	}
	return result
}
```
不过通不过测试，与预期输出有精度上的差异：
```
测试输入
95.123 12
0.4321 20
5.1234 15
6.7592  9
98.999 10
1.0100 12
测试输出
548815620517732158537728
0.000000051485546410769537134231887444413278132060440839268267154693603515625
43992025569.92859649658203125
29448126.7641210220754146575927734375
90429072743629488128
1.1268250301319697737056912956177257001399993896484375
预期输出
548815620517731830194541.899025343415715973535967221869852721
.00000005148554641076956121994511276767154838481760200726351203835429763013462401
43992025569.928573701266488041146654993318703707511666295476720493953024
29448126.764121021618164430206909037173276672
90429072743629540498.107596019456651774561044010001
1.126825030131969720661201
```
尝试将小数点去掉。先计算大整数的幂，再处理小数点：d.go<br>
这次发现跟预期差不多了，不过标准答案有个地方很怪， 0.xxxx非要写成.xxxx~
