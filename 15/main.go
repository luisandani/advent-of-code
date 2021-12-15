package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	grid := loadHeightMap("input")
	walkMap(grid, point{0, 0})
}

func walkMap(grid [][]int, startPnt point) {
	gridHeigth := len(grid)
	gridWidth := len(grid[0])
	destPoint := point{gridWidth - 1, gridHeigth - 1}

	visited := make([][]bool, gridHeigth)
	for i := 0; i < gridHeigth; i++ {
		visited[i] = make([]bool, gridWidth)
	}

	vtx := make(map[point]vertex)
	// we need to create all the possible points
	for row := 0; row < gridHeigth; row++ {
		for col := 0; col < gridWidth; col++ {
			np := point{row, col}
			vtx[np] = vertex{
				dist:       math.MaxInt32,
				prevVertex: nil,
			}
		}
	}

	vtx[startPnt] = vertex{0, &startPnt}

	for {
		currentPoint, err := getNextVisitablePoint(vtx, visited, startPnt)
		if err != nil {
			break
		}

		// moving up
		if isValidPoint(currentPoint.row-1, currentPoint.col, grid, visited) {
			newPoint := point{currentPoint.row - 1, currentPoint.col}
			newDistance := vtx[currentPoint].dist + grid[newPoint.row][newPoint.col]
			if newDistance < vtx[newPoint].dist {
				nvx := vertex{
					dist:       newDistance,
					prevVertex: &currentPoint,
				}
				vtx[newPoint] = nvx
			}
		}

		// moving down
		if isValidPoint(currentPoint.row+1, currentPoint.col, grid, visited) {
			newPoint := point{currentPoint.row + 1, currentPoint.col}
			newDistance := vtx[currentPoint].dist + grid[newPoint.row][newPoint.col]
			if newDistance < vtx[newPoint].dist {
				nvx := vertex{
					dist:       newDistance,
					prevVertex: &currentPoint,
				}
				vtx[newPoint] = nvx
			}
		}

		// moving left
		if isValidPoint(currentPoint.row, currentPoint.col-1, grid, visited) {
			newPoint := point{currentPoint.row, currentPoint.col - 1}
			newDistance := vtx[currentPoint].dist + grid[newPoint.row][newPoint.col]
			if newDistance < vtx[newPoint].dist {
				nvx := vertex{
					dist:       newDistance,
					prevVertex: &currentPoint,
				}
				vtx[newPoint] = nvx
			}
		}

		// moving right
		if isValidPoint(currentPoint.row, currentPoint.col+1, grid, visited) {
			newPoint := point{currentPoint.row, currentPoint.col + 1}
			newDistance := vtx[currentPoint].dist + grid[newPoint.row][newPoint.col]
			if newDistance < vtx[newPoint].dist {
				nvx := vertex{
					dist:       newDistance,
					prevVertex: &currentPoint,
				}
				vtx[newPoint] = nvx
			}
		}

		visited[currentPoint.row][currentPoint.col] = true
	}

	fmt.Printf("Shortest Path to last point: %d\n", vtx[destPoint].dist)
}

func getNextVisitablePoint(vtx map[point]vertex, visited [][]bool, startPoint point) (point, error) {
	if !visited[startPoint.row][startPoint.col] { // if it's the first iteration we return starting point
		return startPoint, nil
	}
	// from the non-visited vortex
	// we need to get the smallest distance that previous was the starting point
	smallestDist := math.MaxInt32
	res := point{-1, -1}
	for p, v := range vtx {
		if p == startPoint || visited[p.row][p.col] { // skip non-pointing to A || already visited
			continue
		}
		if v.dist < smallestDist {
			smallestDist = v.dist
			res = p
		}
	}

	if res.col < 0 || res.row < 0 {
		return res, fmt.Errorf("could not find the next visitable Point")
	}
	return res, nil
}

type vertex struct {
	dist       int
	prevVertex *point
}

func (v vertex) String() string {
	return fmt.Sprintf("Dist:%d From:%v\n", v.dist, v.prevVertex)
}

type point struct {
	row int
	col int
}

func isValidPoint(row int, col int, grid [][]int, visited [][]bool) bool {
	gridHeigth := len(grid)
	gridWidth := len(grid[0])
	if col >= 0 && col < gridWidth && row >= 0 && row < gridHeigth && !visited[row][col] {
		return true
	}
	return false
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
