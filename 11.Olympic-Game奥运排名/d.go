package main

import "fmt"

type record struct {
	Country string
	Golds   int
	Silvers int
	Bronzes int
}

type lessFun func(i, j int) bool

// 在较低的go版本，没有sort.Slice() API可用，自己实现一个冒泡排序
func Sort(records []record, less lessFun) {
	for i := 0; i < len(records); i++ {
		notSwaped := true
		for j := len(records) - 1; j > i; j-- {
			if less(j, j-1) {
				records[j], records[j-1] = records[j-1], records[j]
				notSwaped = false
			}
		}
		if notSwaped {
			break
		}
	}
}

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	records := make([]record, n)
	for i := 0; i < n; i++ {
		var c string
		var g, s, b int
		if _, err := fmt.Scanf("%v %d %d %d", &c, &g, &s, &b); err != nil {
			return
		}
		records[i] = record{Country: c, Golds: g, Silvers: s, Bronzes: b}
	}

	Sort(records, func(i, j int) bool {
		if records[i].Golds > records[j].Golds {
			return true
		}
		if records[i].Golds < records[j].Golds {
			return false
		}
		if records[i].Silvers > records[j].Silvers {
			return true
		}
		if records[i].Silvers < records[j].Silvers {
			return false
		}
		if records[i].Bronzes > records[j].Bronzes {
			return true
		}
		if records[i].Bronzes < records[j].Bronzes {
			return false
		}
		return records[i].Country <= records[j].Country
	})

	for _, r := range records {
		fmt.Println(r.Country)
	}
}
