package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		return
	}
	if n <= 0 || m <= 0 || n > 20 || m > 20 {
		panic("inputed n, m is out of range")
		return
	}

	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	if err != nil {
		return
	}
	word := string(line)
	if len(word) <= 0 {
		panic("inputed word is empty")
		return
	}

	maze := make([]string, n)
	for i := 0; i < n; i++ {
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		maze[i] = string(line)
	}

	if found(maze, word) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

const (
	left = iota
	top
	right
	bottom
)

func found(maze []string, word string) bool {
	n := len(maze)
	m := len(maze[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == word[0] {
				if len(word) == 1 ||
					foundNext(maze, word[1:], i, j, left) ||
					foundNext(maze, word[1:], i, j, top) ||
					foundNext(maze, word[1:], i, j, right) ||
					foundNext(maze, word[1:], i, j, bottom) {
					return true
				}
			}
		}
	}
	return false
}

/*
for maze[i][j]:
left: maze[i][j-1],
top: maze[i-1][j],
right: maze[i][j+1],
bottom: maze[i+1][j]
*/
func foundNext(maze []string, word string, i, j, direct int) bool {
	w := len(word)
	switch direct {
	case left:
		if j-1 < 0 {
			return false
		}
		if maze[i][j-1] == word[0] {
			if w == 1 ||
				// cannot find next at right, because the right charactor has been used
				foundNext(maze, word[1:], i, j-1, left) ||
				foundNext(maze, word[1:], i, j-1, top) ||
				foundNext(maze, word[1:], i, j-1, bottom) {
				return true
			}
		}
		return false
	case top:
		if i-1 < 0 {
			return false
		}
		if maze[i-1][j] == word[0] {
			if w == 1 ||
				// cannot find next at bottom
				foundNext(maze, word[1:], i-1, j, left) ||
				foundNext(maze, word[1:], i-1, j, top) ||
				foundNext(maze, word[1:], i-1, j, right) {
				return true
			}
		}
		return false
	case right:
		if j+1 >= len(maze[0]) {
			return false
		}
		if maze[i][j+1] == word[0] {
			if w == 1 ||
				// cannot find next at left
				foundNext(maze, word[1:], i, j+1, top) ||
				foundNext(maze, word[1:], i, j+1, right) ||
				foundNext(maze, word[1:], i, j+1, bottom) {
				return true
			}
		}
		return false
	case bottom:
		if i+1 >= len(maze) {
			return false
		}
		if maze[i+1][j] == word[0] {
			if w == 1 ||
				// cannot find next at top
				foundNext(maze, word[1:], i+1, j, left) ||
				foundNext(maze, word[1:], i+1, j, right) ||
				foundNext(maze, word[1:], i+1, j, bottom) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
