package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const steps = 40

func main() {
	template, insertRules := loadData("input")
	finalRes := calculateAllPairsPossible(template, insertRules, steps)
	indivCount := getIndividualElementsCount(template, finalRes)
	fmt.Printf("Most - Least Common = %d\n", subtractMostAndLeastCommonElements(indivCount))
}

func subtractMostAndLeastCommonElements(elements map[string]int) int {
	most, less := 0, 0
	for _, v := range elements {
		if most < v {
			most = v
		}
		if less > v || less == 0 {
			less = v
		}
	}
	return most - less
}

func calculateAllPairsPossible(template string, rules map[string]string, steps int) map[string]int {
	mapped := make(map[string]int) //var mapped = mutableMapOf<String, Long>()

	// START - we map the template
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		if i == 0 { // first step we need to insert full string, if not, only the last 2 characters
			mapped[pair] = 1
		} else {
			mapped[pair]++
		}
	}

	// for each step we add/increase the appearances in the map
	for i := 0; i < steps; i++ {
		newMap := make(map[string]int)
		for k, v := range mapped {
			// from NN -> NC and CN
			part1 := k[0:1] + rules[k]
			if _, ok := newMap[part1]; !ok {
				newMap[part1] = v
			} else {
				newMap[part1] += v
			}

			part2 := rules[k] + k[1:2]
			if _, ok := newMap[part2]; !ok {
				newMap[part2] = v
			} else {
				newMap[part2] += v
			}
		}
		mapped = newMap
	}

	return mapped
}

func getIndividualElementsCount(template string, allRes map[string]int) map[string]int {
	charCount := make(map[string]int)
	for k, v := range allRes {
		// we only count the first char
		if _, ok := charCount[k[0:1]]; !ok {
			charCount[k[0:1]] = v
		} else {
			charCount[k[0:1]] += v
		}
	}
	// add the last char from the template
	charCount[template[len(template)-1:]]++
	return charCount
}

func loadData(path string) (template string, insertRules map[string]string) {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template = scanner.Text()
	scanner.Scan() // skip empty line

	insertRules = make(map[string]string)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		// I'm already calculating the replacement
		insertRules[line[0]] = line[1]
	}
	return
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
