package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const mapWidth = 100 // 10 for example (100 for input)

func main() {
	heightMap := loadHeightMap("input")
	mapHeight := len(heightMap)

	basins := make([][]bool, mapHeight)
	for i, _ := range basins {
		basins[i] = make([]bool, mapWidth)
	}

	var lowPoints []int
	var lowPointsCoords []*pointCoord

	for row := 0; row < mapHeight; row++ {
		for col := 0; col < mapWidth; col++ {
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
			if col+1 < mapWidth && heightMap[row][col+1] <= point { // right
				continue
			}
			lowPoints = append(lowPoints, point)
			lowPointsCoords = append(lowPointsCoords, &pointCoord{row, col, 0})
		}
	}

	for _, lpc := range lowPointsCoords {
		lpc.basinSize = calculateIncreaseBasinSize(heightMap, &basins, lpc.x, lpc.y, mapHeight, mapWidth)
	}

	// sort the points based on basinSize
	sort.Slice(lowPointsCoords, func(i, j int) bool {
		return lowPointsCoords[i].basinSize > lowPointsCoords[j].basinSize
	})

	fmt.Printf("Low points sorted: %v\n", lowPointsCoords)
	total := lowPointsCoords[0].basinSize * lowPointsCoords[1].basinSize * lowPointsCoords[2].basinSize
	fmt.Printf("3 Largest basin size multiplied: %d\n", total)
}

func calculateIncreaseBasinSize(heightMap [][]int, basins *[][]bool, row int, col int, mapH int, mapW int) int {
	// take the point and start going moving all the ways until reach a 9
	if (*basins)[row][col] || heightMap[row][col] == 9 {
		return 0
	}

	// if not, we count as
	(*basins)[row][col] = true
	count := 1

	// up, down, left, right
	if row > 0 {
		count += calculateIncreaseBasinSize(heightMap, basins, row-1, col, mapH, mapW)
	}
	if row+1 < mapH {
		count += calculateIncreaseBasinSize(heightMap, basins, row+1, col, mapH, mapW)
	}
	if col > 0 {
		count += calculateIncreaseBasinSize(heightMap, basins, row, col-1, mapH, mapW)
	}
	if col+1 < mapW {
		count += calculateIncreaseBasinSize(heightMap, basins, row, col+1, mapH, mapW)
	}

	return count
}

type pointCoord struct {
	x         int
	y         int
	basinSize int
}

func (p *pointCoord) String() string {
	return fmt.Sprintf("basinSize: %d", p.basinSize)
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
