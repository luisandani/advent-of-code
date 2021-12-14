package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const steps = 10

func main() {
	template, insertRules := loadData("input")

	for i := 0; i < steps; i++ {
		newTemp := ""
		for i := 0; i < len(template)-1; i++ {
			pair := template[i : i+2]
			if i == 0 { // first step we need to insert full string, if not, only the last 2 characters
				newTemp += insertRules[pair]
			} else {
				newTemp += insertRules[pair][1:]
			}

		}
		template = newTemp
	}

	fmt.Printf("Final length:%d\n", len(template))
	fmt.Printf("Most - Least:%d\n", calcMostLeastElementSubstraction(template))
}

func calcMostLeastElementSubstraction(template string) int {
	allElems := make(map[string]int)
	most, least := "", ""
	for _, v := range template {
		allElems[string(v)]++
		if allElems[string(v)] > allElems[most] {
			most = string(v)
		}
		if allElems[string(v)] < allElems[least] || allElems[least] == 0 {
			least = string(v)
		}
	}
	fmt.Printf("%+v\n", allElems)
	return allElems[most] - allElems[least]
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
		insertRules[line[0]] = fmt.Sprintf("%s%s%s", line[0][0:1], line[1], line[0][1:])
	}
	return
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
