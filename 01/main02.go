package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	var depthList []int
	// let's read all the lines
	for scan.Scan() {
		newDepth, err := strconv.Atoi(scan.Text())
		if err != nil {
			fmt.Printf("Error converting text to int. %s\n", err)
		}

		depthList = append(depthList, newDepth)
	}
	if err := scan.Err(); err != nil {
		panic(fmt.Sprintf("Scanner contained an error: %s\n", err))
	}

	var queue []int
	counter := 0

	for _, v := range depthList {
		queue = append(queue, v)
		if len(queue) == 4 {
			a := queue[0] + queue[1] + queue[2]
			b := queue[1] + queue[2] + queue[3]
			if b > a {
				counter++
			}
			queue = queue[1:4]
		}
	}
	// how many?
	fmt.Printf("### %d measurements larger than the previous one.\n", counter)
}
