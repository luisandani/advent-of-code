package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	points, foldInstructions := loadInputData("input")
	fmt.Printf("instructions: %v\n\n", foldInstructions)

	newPoints := make(map[point]struct{})

	maxWidth, maxHeigth := 0, 0

	for _, p := range points {
		finalPoint := point{col: p.col, row: p.row}
		for _, inst := range foldInstructions {
			transfType := strings.Split(inst, "=")[0]
			split, err := strconv.Atoi(strings.Split(inst, "=")[1])
			must(err)

			switch transfType {
			case "y":
				if finalPoint.row > split {
					finalPoint = point{
						col: finalPoint.col,
						row: 2*split - finalPoint.row,
					}
				}
			case "x":
				if finalPoint.col > split {
					finalPoint = point{
						col: 2*split - finalPoint.col,
						row: finalPoint.row,
					}
				}
			}
		}
		newPoints[finalPoint] = struct{}{}
		if maxHeigth < finalPoint.row {
			maxHeigth = finalPoint.row
		}
		if maxWidth < finalPoint.col {
			maxWidth = finalPoint.col
		}
	}

	fmt.Printf("%+v\nTotal: %d \n", newPoints, len(newPoints))

	for row := 0; row <= maxHeigth; row++ {
		for col := 0; col <= maxWidth; col++ {
			coord := point{row: row, col: col}
			if _, ok := newPoints[coord]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func loadInputData(path string) ([]point, []string) { //([][]bool, []string) {
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

	return allPoints, foldInstructions
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
