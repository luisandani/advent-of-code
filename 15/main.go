package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	grid := loadHeightMap("example")
	fmt.Printf("%v\n", grid)
	walkMap(grid, point{0, 0})
}

func walkMap(grid [][]int, startPnt point) {
	gridHeigth := len(grid)
	gridWidth := len(grid[0])
	destPoint := point{gridWidth - 1, gridHeigth - 1}
	fmt.Printf("start: %v | dest: %v\n", startPnt, destPoint)

	visited := make([][]bool, gridHeigth)
	for i := 0; i < gridHeigth; i++ {
		visited[i] = make([]bool, gridWidth)
	}
	visited[startPnt.row][startPnt.col] = true

	shortestPath := 0

	// applying BFS on matrix cells starting from source
	var queue []qItem
	queue = append(queue, qItem{row: startPnt.col, col: startPnt.row, dist: 0})
	for {
		currentPoint, queue := queue[len(queue)-1], queue[:len(queue)-1]
		fmt.Printf("Visited: %v\n", currentPoint)
		time.Sleep(500 * time.Millisecond)
		// Destination found;
		if currentPoint.row == destPoint.row && currentPoint.col == destPoint.col {
			shortestPath = currentPoint.dist
			break
		}

		// moving up
		if isValidPoint(currentPoint.row-1, currentPoint.col, grid, visited) {
			nr := currentPoint.row - 1
			nc := currentPoint.col
			nd := currentPoint.dist + grid[nr][nc]
			queue = append(queue, qItem{nr, nc, nd})
			visited[nr][nc] = true
		}

		// moving down
		if isValidPoint(currentPoint.row+1, currentPoint.col, grid, visited) {
			nr := currentPoint.row + 1
			nc := currentPoint.col
			nd := currentPoint.dist + grid[nr][nc]
			queue = append(queue, qItem{nr, nc, nd})
			visited[nr][nc] = true
		}

		// moving left
		if isValidPoint(currentPoint.row, currentPoint.col-1, grid, visited) {
			nr := currentPoint.row
			nc := currentPoint.col - 1
			nd := currentPoint.dist + grid[nr][nc]
			queue = append(queue, qItem{nr, nc, nd})
			visited[nr][nc] = true
		}

		// moving right
		if isValidPoint(currentPoint.row, currentPoint.col+1, grid, visited) {
			nr := currentPoint.row
			nc := currentPoint.col + 1
			nd := currentPoint.dist + grid[nr][nc]
			queue = append(queue, qItem{nr, nc, nd})
			visited[nr][nc] = true
		}

		if len(queue) < 1 {
			break
		}
	}

	fmt.Printf("shortest path: %d", shortestPath)
}

func isValidPoint(row int, col int, grid [][]int, visited [][]bool) bool {
	gridHeigth := len(grid)
	gridWidth := len(grid[0])
	if col >= 0 && col < gridWidth && row >= 0 && row < gridHeigth && !visited[row][col] {
		return true
	}
	return false
}

// qItem for current location and distance from source location
type qItem struct {
	row  int
	col  int
	dist int
}

type point struct {
	row int
	col int
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
