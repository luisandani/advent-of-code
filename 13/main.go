package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//const gridWidth, gridHeight = 11, 15

func main() {
	paperGrid, foldInstructions := loadInputData("example")
	printGrid(paperGrid)
	fmt.Printf("instructions: %v\n\n", foldInstructions)

	for _, instr := range foldInstructions {
		axis := strings.Split(instr, "=")[0]
		coord, err := strconv.Atoi(strings.Split(instr, "=")[1])
		must(err)
		switch axis {
		case "x":
			paperGrid = foldByColumn(paperGrid, coord)
		case "y":
			paperGrid = foldByRow(paperGrid, coord)
		default:
			panic("this fold type must not happen")
		}
		printGrid(paperGrid) // comment for big one
		break
	}
	printGrid(paperGrid)
}

func foldByColumn(originalGrid [][]bool, coord int) [][]bool {
	fmt.Printf("folding X axis, coord %d\n", coord)
	newGridColumns := calcNewAmountOfCols(originalGrid, coord)
	newGridRows := len(originalGrid)

	var resGrid [][]bool
	for row := 0; row < newGridRows; row++ {
		newRow := make([]bool, newGridColumns)
		for col := 0; col < newGridColumns; col++ {
			if originalGrid[row][col] || originalGrid[row][coord*2-col-1] {
				newRow[col] = true
			}
		}
		resGrid = append(resGrid, newRow)
	}
	return resGrid
}

func calcNewAmountOfCols(grid [][]bool, coord int) int {
	// we need to get the longest part after the fold
	if len(grid[0])-coord > coord {
		return len(grid[0]) - coord
	}
	return coord
}

func foldByRow(originalGrid [][]bool, coord int) [][]bool {
	fmt.Printf("folding Y axis, coord %d\n", coord)
	newGridColumns := len(originalGrid[0]) // stays the same
	newGridRows := calcNewAmountOfRows(originalGrid, coord)

	var resGrid [][]bool
	for row := 0; row < newGridRows; row++ {
		newRow := make([]bool, newGridColumns)
		for col := 0; col < newGridColumns; col++ {
			if originalGrid[row][col] || originalGrid[row][coord*2-col-1] {
				newRow[col] = true
			}
		}
		resGrid = append(resGrid, newRow)
	}
	
	return resGrid
}

func calcNewAmountOfRows(grid [][]bool, coord int) int {
	// we need to get the longest part after the fold
	if len(grid)-coord-1 > coord {
		return len(grid) - coord - 1
	}
	return coord
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, col := range row {
			if col {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}
}

func loadInputData(path string) ([][]bool, []string) {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var allPoints []point
	var foldInstructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if strings.HasPrefix(line, "fold along") {
			foldInstructions = append(foldInstructions, strings.Replace(line, "fold along ", "", 1))
		} else {
			coords := strings.Split(line, ",")
			col, err := strconv.Atoi(coords[0])
			must(err)
			row, err := strconv.Atoi(coords[1])
			must(err)
			//paperGrid[col][row] = true
			allPoints = append(allPoints, point{row, col})
		}
	}
	maxHeight, maxWidth := 0, 0
	for _, p := range allPoints {
		if p.col+1 > maxWidth {
			maxWidth = p.col + 1
		}
		if p.row+1 > maxHeight {
			maxHeight = p.row + 1
		}
	}

	paperGrid := make([][]bool, maxHeight)
	for i := 0; i < maxHeight; i++ {
		paperGrid[i] = make([]bool, maxWidth)
	}
	for _, p := range allPoints {
		paperGrid[p.row][p.col] = true
	}

	return paperGrid, foldInstructions
}

type point struct {
	row int
	col int
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
