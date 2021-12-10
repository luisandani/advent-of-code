package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []rune
		for _, v := range scanner.Text() {
			line = append(line, v)
		}
		input = append(input, line)
	}

	var illegalChars []rune
	for _, l := range input {
		var queue []rune
		for _, r := range l {
			syntaxErr := false
			switch r {
			case '(', '[', '{', '<':
				queue = append(queue, r)
			case ')':
				if queue[len(queue)-1] == '(' {
					queue = queue[:len(queue)-1] // pop
				} else {
					syntaxErr = true
					illegalChars = append(illegalChars, r)
				}
			case ']':
				if queue[len(queue)-1] == '[' {
					queue = queue[:len(queue)-1] // pop
				} else {
					syntaxErr = true
					illegalChars = append(illegalChars, r)
				}
			case '}':
				if queue[len(queue)-1] == '{' {
					queue = queue[:len(queue)-1] // pop
				} else {
					syntaxErr = true
					illegalChars = append(illegalChars, r)
				}
			case '>':
				if queue[len(queue)-1] == '<' {
					queue = queue[:len(queue)-1] // pop
				} else {
					syntaxErr = true
					illegalChars = append(illegalChars, r)
				}
			}

			if syntaxErr {
				break
			}
		}
	}

	syntaxErrorPoints := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	points := 0
	for _, r := range illegalChars {
		points += syntaxErrorPoints[r]
	}
	fmt.Printf("IllegalChars: %s\nPoints: %d\n", string(illegalChars), points)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
