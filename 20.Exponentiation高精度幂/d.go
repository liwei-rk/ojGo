package main

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

type input struct {
	r string
	n int
}

func main() {
	inputs := make([]input, 0)
	for {
		var r string
		var n int
		if _, err := fmt.Scanf("%v %d", &r, &n); err != nil {
			break
		}
		inputs = append(inputs, input{r: r, n: n})
	}
	for _, input := range inputs {
		fmt.Println(power(input.r, input.n))
	}
}

func power(r string, n int) string {
	const dot = "."
	dotIndex := strings.Index(r, dot)
	decimalBits := 0
	if dotIndex != -1 { // has dot
		r = strings.TrimRight(r, "0")
		decimalBits = len(r) - dotIndex - 1
	}
	r = strings.Replace(r, dot, "", -1)
	bigR := big.NewInt(0)
	if _, ok := bigR.SetString(r, 10); !ok {
		panic("invalid input")
	}
	bigResult := big.NewInt(0)
	bigResult.Add(bigResult, bigR)
	for i := 1; i < n; i++ {
		bigResult.Mul(bigResult, bigR)
	}

	result := bigResult.String()
	if decimalBits > 0 {
		dotIndex := len(result) - decimalBits*n
		if dotIndex >= 0 {
			result = result[:dotIndex] + dot + result[dotIndex:]
		} else {
			prefix := bytes.Buffer{}
			prefix.WriteString("0.")
			for i := 0; i < -dotIndex; i++ {
				prefix.WriteString("0")
			}
			result = prefix.String() + result
		}
	}

	if strings.HasPrefix(result, "0.") { // 标准答案预期要消去小数点前边的0，妥协一下~~~~
		result = result[1:]
	}
	return result
}
