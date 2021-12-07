package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const gridSize = 1000

func main() {
	ventLines := loadInput("input")
	var diagram [gridSize][gridSize]int
	for _, vl := range ventLines {
		if vl.from.x != vl.to.x && vl.from.y != vl.to.y {
			continue
		}
		diagram = addLineToDiagram(diagram, vl)
	}
	//printGrid(diagram)
	fmt.Printf("Overlaps: %d\n", calcOverlaps(diagram))
}

func calcOverlaps(d [gridSize][gridSize]int) int {
	overlaps := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if d[x][y] > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func addLineToDiagram(d [gridSize][gridSize]int, l line) [gridSize][gridSize]int {
	if l.from.y == l.to.y { // run on the x
		steps, direction := getStepsAndDirection(l.from.x, l.to.x)
		for i := 0; i <= steps; i++ {
			d[l.from.x+(i*direction)][l.from.y]++
		}
	} else if l.from.x == l.to.x { // run on the y
		steps, direction := getStepsAndDirection(l.from.y, l.to.y)
		for i := 0; i <= steps; i++ {
			d[l.from.x][l.from.y+(i*direction)]++
		}
	}
	return d
}

func getStepsAndDirection(from int, to int) (int, int) {
	steps := to - from
	direction := 1
	if steps < 0 {
		direction = -1
		steps = int(math.Abs(float64(steps)))
	}
	return steps, direction
}

func loadInput(path string) []line {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var ventLines []line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ventLines = append(ventLines, readLine(scanner.Text()))
	}

	return ventLines
}

func readLine(text string) line {
	points := strings.Split(text, " -> ")
	ft := strings.Split(points[0], ",")
	tt := strings.Split(points[1], ",")
	return line{NewPoint(ft[0], ft[1]), NewPoint(tt[0], tt[1])}
}

type line struct {
	from point
	to   point
}

type point struct {
	x int
	y int
}

func NewPoint(xs string, ys string) point {
	x, err := strconv.Atoi(xs)
	must(err)
	y, err := strconv.Atoi(ys)
	must(err)
	return point{x, y}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// printGrid only useful for example
//func printGrid(grid [gridSize][gridSize]int) {
//	for i := 0; i < gridSize; i++ {
//		for j := 0; j < gridSize; j++ {
//			v := grid[j][i]
//			s := "."
//			if v > 0 {
//				s = strconv.Itoa(v)
//			}
//			fmt.Printf("%s", s)
//		}
//		fmt.Println("")
//	}
//	fmt.Println("")
//}
