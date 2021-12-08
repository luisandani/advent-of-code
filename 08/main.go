package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	//var signals [][]string
	digAppear := map[int]int{2: 0, 3: 0, 4: 0, 7: 0}
	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 10 patterns + | (delimiter) + 4 digit output value
		_, out := getPatternsAndOutputFromline(scanner.Text())
		for _, o := range out {
			sc := countSegments(o)
			if _, ok := digAppear[sc]; ok {
				digAppear[sc]++
				counter++
			}
		}
	}
	fmt.Printf("Total times: %d | segments appearance: %v\n", counter, digAppear)
	// unique numbers:
	// 		1: uses 2 segments
	// 		4: uses 4 segments
	// 		7: uses 3 segments
	// 		8: uses 7 (all) segments
	// 		rest: uses 5-6 segments

}

func getPatternsAndOutputFromline(textLine string) (patterns []string, output []string) {
	r := strings.Split(textLine, " | ")
	return strings.Split(r[0], " "), strings.Split(r[1], " ")
}

func countSegments(signal string) int {
	segments := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0}
	count := 0
	for _, v := range signal {
		s := fmt.Sprintf("%c", v)
		segments[s]++
		if segments[s] == 1 { // first time we add, we increase the counter
			count++
		}
	}
	return count
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
