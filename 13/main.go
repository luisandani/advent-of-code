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
	split := 655
	for _, p := range points {
		np := point{col: p.col, row: p.row}
		if p.col > split { //if p.row > split {
			np = point{
				col: 2*split - p.col,
				row: p.row, //row: 2*split - p.row,
			}
		}
		newPoints[np] = struct{}{}
	}
	fmt.Printf("Total: %d\n", len(newPoints))
}

func loadInputData(path string) ([]point, []string) {
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
