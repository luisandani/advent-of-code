package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var count int = 0

func main() {
	caveMap := loadMap("input")
	walkPath(caveMap, []string{}, "start")
	fmt.Printf("Total paths: %d\n", count)
}

func walkPath(caveMap map[string][]string, visitedCaves []string, currentCave string) {
	// add current cave
	visitedCaves = append(visitedCaves, currentCave)
	// if we reached the end, stop
	if currentCave == "end" {
		//fmt.Printf("End reached: %v\n", visitedCaves)
		count++
		return
	}
	// we go through all possible paths
	for _, nc := range caveMap[currentCave] {
		if isVisitable(visitedCaves, nc) {
			walkPath(caveMap, visitedCaves, nc)
		}
	}
}

func isVisitable(visitedCaves []string, cave string) bool {
	for _, c := range visitedCaves {
		if !isBigCave(cave) && c == cave {
			return false
		}
	}
	return true
}

func isBigCave(cave string) bool {
	if cave == strings.ToUpper(cave) {
		return true
	}
	return false
}

func loadMap(path string) map[string][]string {
	caveMap := make(map[string][]string)
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		caveMap[line[0]] = append(caveMap[line[0]], line[1])
		caveMap[line[1]] = append(caveMap[line[1]], line[0])
	}

	return caveMap
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
