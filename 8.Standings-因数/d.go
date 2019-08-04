package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	var n uint64
	for {
		if _, err := fmt.Scanf("%d", &n); err != nil {
			return
		}
		if n == 0 {
			return
		}
		r1, r2 := Factors(n)
		fmt.Print(r1.Len() + r2.Len())
		for item := r1.Front(); item != nil; item = item.Next() {
			fmt.Printf(" %v", item.Value)
		}
		for item := r2.Front(); item != nil; item = item.Next() {
			fmt.Printf(" %v", item.Value)
		}
		fmt.Println()
	}
}

// 两个列表，只需要遍历到sqrt(n)
// 如果用一个数组，需要遍历到n
func Factors(n uint64) (*list.List, *list.List) {
	r1 := list.New()
	r2 := list.New()
	last := uint64(math.Sqrt(float64(n)))
	for i := uint64(1); i <= last; i++ {
		if n%i == 0 {
			r1.PushBack(i)

			j := n / i
			if j != i {
				r2.PushFront(j)
			}
		}
	}
	return r1, r2
}
