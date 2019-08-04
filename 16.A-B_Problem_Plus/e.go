package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
	"strings"
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
	s = strings.TrimLeft(s, "0")
	if s == "" {
		s = "0"
	}
	r := &BigInt{str: s}
	for _, c := range s {
		r.vals.PushFront(int(c - '0'))
	}
	return r
}

func (a *BigInt) Sub(b *BigInt) string {
	var nodeA, nodeB *list.Element
	sign := true // true: + , false: -
	if a.IsMoreThan(b) {
		nodeA = a.vals.Front()
		nodeB = b.vals.Front()
	} else {
		nodeB = a.vals.Front()
		nodeA = b.vals.Front()
		sign = false
	}
	r := list.New()
	lend := 0 // must be 0 or -1
	for nodeA != nil {
		var a, b, sub int
		if nodeA != nil {
			a = nodeA.Value.(int)
			nodeA = nodeA.Next()
		}
		if nodeB != nil {
			b = nodeB.Value.(int)
			nodeB = nodeB.Next()
		}
		a += lend
		sub = a - b
		if sub < 0 {
			lend = -1
			sub = 10 + a - b
		} else {
			lend = 0
		}
		r.PushFront(sub % 10)
	}

	buffer := bytes.Buffer{}
	for node := r.Front(); node != nil; node = node.Next() {
		val := node.Value.(int)
		s := strconv.Itoa(val)
		buffer.WriteString(s)
	}

	s := buffer.String()
	s = strings.TrimLeft(s, "0")
	if s == "" {
		s = "0"
	}
	if !sign && s != "0" {
		return "-" + s
	}
	return s
}

func (a *BigInt) IsMoreThan(b *BigInt) bool {
	if a == nil || b == nil {
		return false
	}
	if a.vals.Len() > b.vals.Len() {
		return true
	}
	if a.vals.Len() < b.vals.Len() {
		return false
	}
	nodeA := a.vals.Back()
	nodeB := b.vals.Back()
	for nodeA != nil {
		aVal := nodeA.Value.(int)
		bVal := nodeB.Value.(int)
		if aVal > bVal {
			return true
		}
		if aVal < bVal {
			return false
		}
		nodeA = nodeA.Prev()
		nodeB = nodeB.Prev()
	}
	return false
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
		fmt.Printf("%s - %s = %s\n\n", c.a.str, c.b.str, c.a.Sub(c.b))
	}
}
