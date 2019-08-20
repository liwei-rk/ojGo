## 动态规划涂砖块
```
题目描述
在美丽的尧山，有一个大广场，50周年校庆的时候Solo就在大广场上见证了史上最壮观的焰火。
在广场上有一排方砖是有颜色的，被涂上红色或者绿色，从左到右排列。
现在校方要求重新喷涂颜色，但不一定要每一块方砖都重新喷涂，
因为校方的目的是：每一块红色的方砖都至少在绿色方砖的左边（也就是每一个红的左边不能有绿的），
并且尽量喷涂最少的次数。

解答要求
时间限制：1000ms, 内存限制：64MB
输入
输入只有一行，包含一个字符串S，且只包含'R'(代表红色)或者'G'(代表绿色)。
我们保证字符串S的长度L的范围是(0 < L < 50 )。

输出
输出需要重新喷涂的方砖的最少数量。

样例
输入样例 1

RGRGR
输出样例 1

2

提示
样例中字符串S为RGRGR，则我们可以这么喷涂，
即RGRGR喷成RRRGG（即将第二个字符喷成R，最后一个字符喷成G）
或者RRRRR（即将两个G都喷成R），都是只需喷涂两个方砖，所以答案为2。

我们再举个例子:若S为RRRGGGGG，
则我们不需要在重新喷涂就已经满足“每一块红色的方砖都在绿色方砖的左边”的要求，所以答案将是0。
```
```
对于第i个砖块，要吗喷成红色，要吗喷成绿色
我们用3个长度为n的一维数组paintedReds，paintedGreens和dp
其中dp[i]表示走到第i步需要喷涂的最少砖块数
paintedReds[i]表示走到第i步时绿色喷成红色的砖块数
如果第i块砖是红色，那么不用喷，所以paintedReds[i] = paintedReds[i-1]
如果第i块砖是绿色，那么需要喷，所以paintedReds[i] = paintedReds[i-1] + 1
paintedGreens[i]表示走到第i步时，红色转为绿色的砖块数
如果第i块砖是红色，那么需要喷，所以paintedGreens[i] = dp[i-1] + 1
如果第i块砖是绿色，那么不用喷，所以paintedGreens[i] = dp[i-1]
注意红色喷成绿色跟之前的状态无关，所以利用之前最好的结果来计数
而dp[i] = min(paintedReds[i], paintedGreens[i])
```
实现如下：
```
func leastPaintBricks(s string) int {
	const red = 'R'
	n := len(s)
	dp := make([]int, n)
	paintedReds := make([]int, n)
	paintedGreens := make([]int, n)
	/*
	dp[i]=min(paintedReds[i], paintedGreens[i])
	where paintedGreens[i]=dp[i-1]+s[i]==‘R’?1:0
	paintedReds[i]=paintedReds[i-1]+s[i]==‘R’?0:1
	*/
	if s[0] == red {
		paintedReds[0] = 0
		paintedGreens[0] = 1
	} else {
		paintedReds[0] = 1
		paintedGreens[0] = 0
	}
	dp[0] = 0 // dp[i] = min(paintedReds[i], paintedGreens[i])
	for i := 1; i < n; i++ {
		if s[i] == red {
			paintedReds[i] = paintedReds[i-1]
			paintedGreens[i] = dp[i-1] + 1
		} else {
			paintedGreens[i] = dp[i-1]
			paintedReds[i] = paintedReds[i-1] + 1
		}
		if paintedReds[i] < paintedGreens[i] {
			dp[i] = paintedReds[i]
		} else {
			dp[i] = paintedGreens[i]
		}
	}
	return dp[n-1]
}
```
还可以有一个小小的优化，paintedGreens数组是可以节省的，只需要dp和paintedReds两个数组即可：
```
func leastPaintBricks(s string) int {
	const red = 'R'
	n := len(s)
	dp := make([]int, n)
	paintedReds := make([]int, n)
	if s[0] != red {
		paintedReds[0] = 1
	}

	for i := 1; i < n; i++ {
		if s[i] == red {
			paintedReds[i] = paintedReds[i-1]
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
			paintedReds[i] = paintedReds[i-1] + 1
		}
		// dp[i] = min(paintedReds[i], dp[i])
		if paintedReds[i] < dp[i] {
			dp[i] = paintedReds[i]
		}
	}
	return dp[n-1]
}
```
