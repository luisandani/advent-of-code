package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	input := readInputFromFile("input")
	incompleteLines := getIncompleteLines(input)

	var allPoints []int
	for _, v := range incompleteLines {
		var queue []rune
		for _, v := range v {
			// opening
			if strings.ContainsRune("([{<", v) {
				queue = append(queue, v)
				continue
			}
			// closing
			switch v {
			case ')':
				if queue[len(queue)-1] == '(' {
					queue = queue[:len(queue)-1] // pop
				}
			case ']':
				if queue[len(queue)-1] == '[' {
					queue = queue[:len(queue)-1] // pop
				}
			case '}':
				if queue[len(queue)-1] == '{' {
					queue = queue[:len(queue)-1] // pop
				}
			case '>':
				if queue[len(queue)-1] == '<' {
					queue = queue[:len(queue)-1] // pop
				}
			default:
				panic("that should not happen")
			}
		}
		missingRunes := revertSlice(queue)
		points := map[rune]int{
			')': 1,
			']': 2,
			'}': 3,
			'>': 4,
		}
		totalPoints := 0
		for _, r := range missingRunes {
			totalPoints = totalPoints*5 + points[r]
		}
		allPoints = append(allPoints, totalPoints)
	}
	sort.Ints(allPoints)
	fmt.Printf("middle score: %d\n", allPoints[len(allPoints)/2])
}

func revertSlice(queue []rune) []rune {
	var rev []rune
	for i := len(queue) - 1; i >= 0; i-- {
		rev = append(rev, getClosingRune(queue[i]))
	}
	return rev
}

func getClosingRune(orig rune) rune {
	switch orig {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return '#'
}

func getIncompleteLines(input [][]rune) [][]rune {
	var incompleteLines [][]rune
	for _, l := range input {
		var queue []rune
		corrupted := false
		for _, v := range l {
			switch v {
			case '(', '[', '{', '<':
				queue = append(queue, v)
			case ')':
				if queue[len(queue)-1] == '(' {
					queue = queue[:len(queue)-1] // pop
				} else {
					corrupted = true
				}
			case ']':
				if queue[len(queue)-1] == '[' {
					queue = queue[:len(queue)-1] // pop
				} else {
					corrupted = true
				}
			case '}':
				if queue[len(queue)-1] == '{' {
					queue = queue[:len(queue)-1] // pop
				} else {
					corrupted = true
				}
			case '>':
				if queue[len(queue)-1] == '<' {
					queue = queue[:len(queue)-1] // pop
				} else {
					corrupted = true
				}
			}
			if corrupted {
				break
			}
		}
		if !corrupted {
			incompleteLines = append(incompleteLines, l)
		}
	}
	return incompleteLines
}

func readInputFromFile(path string) [][]rune {
	file, err := os.Open(path)
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
	return input
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
