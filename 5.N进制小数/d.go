package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var m float64
	var n int
	for {
		if _, err := fmt.Scanf("%v %d", &m, &n); err != nil {
			return
		}
		if math.Abs(m-0.0) < 0.000001 && n <= 0 {
			return
		}
		fmt.Println(parse(m, n, 10))
	}
}

/*
十进制整数转换成n进制整数，用不断取模后模的逆序获得 —— strconv包已经实现
十进制小数数转换成n进制小数，就是小数部分不断乘以n取其整数部分，整数的顺序序列即为所求的小数

decimal: 要转换的十进制小数
n： 要转换的进制 n<10
bits： 要保留的精度
*/
func parse(decimal float64, base, precision int) string {
	i := int64(decimal) //整数部分
	buffer := bytes.Buffer{}
	buffer.WriteString(strconv.FormatInt(i, base))
	buffer.WriteString(".")
	for i := 0; i < precision; i++ {
		decimal *= float64(base)
		c := int(decimal) //整数部分
		decimal -= float64(c)
		buffer.WriteString(strconv.Itoa(c))
	}
	return buffer.String()
}
