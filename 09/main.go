package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mapLength = 100 // 10 for example (100 for input)

func main() {
	heightMap := loadHeightMap("input")

	var lowPoints []int
	mapHeight := len(heightMap)
	for row := 0; row < mapHeight; row++ {
		for col := 0; col < mapLength; col++ {
			point := heightMap[row][col]
			// check up, down, left, right. if only in 1 is already higher, move to the next one
			if row > 0 && heightMap[row-1][col] <= point { // up
				continue
			}
			if row+1 < mapHeight && heightMap[row+1][col] <= point { // down
				continue
			}
			if col > 0 && heightMap[row][col-1] <= point { // left
				continue
			}
			if col+1 < mapLength && heightMap[row][col+1] <= point { // right
				continue
			}
			lowPoints = append(lowPoints, point)
		}
	}

	fmt.Printf("Risk: %d\n", calculateRisk(lowPoints))
}

func calculateRisk(lowPoints []int) int {
	res := 0
	for _, point := range lowPoints {
		res += point + 1
	}
	return res
}

func loadHeightMap(path string) [][]int {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var heightMap [][]int
	for scanner.Scan() {
		var newRow []int
		for _, v := range scanner.Text() {
			h, err := strconv.Atoi(string(v))
			must(err)
			newRow = append(newRow, h)
		}
		heightMap = append(heightMap, newRow)
	}
	return heightMap
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
