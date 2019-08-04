package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type BigInt struct {
	vals list.List
	str  string
}

type eCase struct {
	a *BigInt
	b *BigInt
}

func New(s string) *BigInt {
	r := &BigInt{str: s}
	for _, c := range s {
		r.vals.PushFront(int(c - '0'))
	}
	return r
}

func (a *BigInt) Add(b *BigInt) string {
	nodeA := a.vals.Front()
	nodeB := b.vals.Front()
	r := list.New()
	carry := 0 // must be 0 or 1
	for nodeA != nil || nodeB != nil || carry != 0 {
		var a, b, sum int
		if nodeA != nil {
			a = nodeA.Value.(int)
			nodeA = nodeA.Next()
		}
		if nodeB != nil {
			b = nodeB.Value.(int)
			nodeB = nodeB.Next()
		}
		sum = a + b + carry
		r.PushFront(sum % 10)
		carry = sum / 10
	}

	buffer := bytes.Buffer{}
	for node := r.Front(); node != nil; node = node.Next() {
		s := strconv.Itoa(node.Value.(int))
		buffer.WriteString(s)
	}
	return buffer.String()
}

func main() {
	var t int
	if _, err := fmt.Scanf("%d", &t); err != nil {
		return
	}
	if t < 1 || t > 20 {
		return
	}
	var cases = make([]eCase, t)
	for i := 0; i < t; i++ {
		var a, b string
		if _, err := fmt.Scanf("%v %v", &a, &b); err != nil {
			panic("invalid input")
		}
		if len(a) == 0 || len(b) == 0 {
			panic("invalid input, number cann't be empty")
		}
		if a[0] == '-' || b[0] == '-' {
			panic("invalid input, number must be positive")
		}
		cases[i].a = New(a)
		cases[i].b = New(b)
	}

	for i, c := range cases {
		fmt.Printf("Case %d:\n", i+1)
		fmt.Printf("%s + %s = %s\n\n", c.a.str, c.b.str, c.a.Add(c.b))
	}
}
