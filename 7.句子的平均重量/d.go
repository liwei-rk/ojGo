package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		r := bufio.NewReader(os.Stdin)
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		s := string(line)
		if len(s) == 0 {
			return
		}
		words := strings.Fields(s)
		totalLen := 0
		for _, word := range words {
			totalLen += len(word)
		}
		fmt.Printf("%.2f\n", float64(totalLen)/float64(len(words)))
	}
}
