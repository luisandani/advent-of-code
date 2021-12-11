package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const gridSize = 10

var octoGrid [][]int
var flashedGrid [][]bool
var currentStep int

func main() {
	octoGrid = loadGrid("input")
	flashedGrid = cleanFlashGrid()

	for s := 0; s < 1000; s++ {
		currentStep = s
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				triggerOctopus(i, j)
			}
		}
		flashedGrid = cleanFlashGrid()
	}
}

func triggerOctopus(i int, j int) {
	if flashedGrid[i][j] { // if already flashed
		return
	}

	// increase 1 the level
	octoGrid[i][j]++
	// if >9 flash and trigger adjacent octos
	if octoGrid[i][j] > 9 {
		octoGrid[i][j] = 0
		flashedGrid[i][j] = true
		didAllOctopusFlashed()
		if i > 0 && j > 0 { // up-left
			triggerOctopus(i-1, j-1)
		}
		if j > 0 { // up
			triggerOctopus(i, j-1)
		}
		if i < gridSize-1 && j > 0 { // up-right
			triggerOctopus(i+1, j-1)
		}
		if i > 0 && j < gridSize-1 { // down-left
			triggerOctopus(i-1, j+1)
		}
		if j < gridSize-1 { // down
			triggerOctopus(i, j+1)
		}
		if i < gridSize-1 && j < gridSize-1 { // down-right
			triggerOctopus(i+1, j+1)
		}
		if i > 0 { // left
			triggerOctopus(i-1, j)
		}
		if i < gridSize-1 { // right
			triggerOctopus(i+1, j)
		}

	}
}

func didAllOctopusFlashed() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if flashedGrid[i][j] == false {
				return
			}
		}
	}
	fmt.Printf("all octopus flashed at %d\n", currentStep+1)
	os.Exit(1)
}

func sumFlashes(grid [][]bool) int {
	res := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] {
				res++
			}
		}
	}
	return res
}

func loadGrid(path string) [][]int {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var res [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var nl []int
		for _, v := range scanner.Text() {
			i, err := strconv.Atoi(string(v))
			must(err)
			nl = append(nl, i)
		}
		res = append(res, nl)
	}
	return res
}

func cleanFlashGrid() [][]bool {
	res := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		res[i] = make([]bool, gridSize)
	}
	return res
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
