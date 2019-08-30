package main

import "fmt"

type fruit struct {
	id    int
	score int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n < 1 || n > 3000 {
		panic("input out of range")
	}
	inputs := make([]fruit, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &inputs[i].id, &inputs[i].score)
		if inputs[i].id < 1 || inputs[i].id > 10 || inputs[i].score < 1 || inputs[i].score > 20 {
			panic("invalid input")
		}
	}
	score, count := act(inputs)
	fmt.Println(score, count)
}

func act(fruits []fruit) (int, int) {
	count := 0
	n := len(fruits)
	if n == 0 {
		return 0, count
	}
	maxScore := fruits[0].score
	count ++
	lastCutedIndex := 0
	for i := 1; i < n; i++ {
		switch {
		case fruits[i].id != fruits[i-1].id:
			count ++
			lastCutedIndex = i
			maxScore += fruits[i].score
		case fruits[i].score > 10:
			count ++
			lastCutedIndex = i
			maxScore += fruits[i].score - 10
		case fruits[lastCutedIndex].score < fruits[i].score: // not cut fruit[lastCutedIndex] but cut the current one
			lastCutedIndex = i
			maxScore += fruits[i].score - fruits[lastCutedIndex].score
		}
	}
	return maxScore, count
}
