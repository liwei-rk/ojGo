package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		if len(line) == 0 {
			return
		}
		fmt.Println(soloSentence(string(line)))
	}
}

func soloSentence(s string) string {
	dis := 'a' - 'A'
	buffer := bytes.Buffer{}
	for _, c := range s {
		switch {
		case c == 'A' || c == 'O' || c == 'E' || c == 'I' || c == 'U':
			buffer.WriteRune(c)
		case c == 'a' || c == 'o' || c == 'e' || c == 'i' || c == 'u':
			buffer.WriteRune(c - dis)
		case c >= 'A' && c <= 'Z':
			buffer.WriteRune(c + dis)
		default:
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
