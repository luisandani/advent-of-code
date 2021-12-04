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

	var queue []int
	counter := 0
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		newDepth, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(err)
		}

		queue = append(queue, newDepth)
		if len(queue) == 4 {
			a := queue[0] + queue[1] + queue[2]
			b := queue[1] + queue[2] + queue[3]
			if b > a {
				counter++
			}
			queue = queue[1:4]
		}

	}
	if err := scan.Err(); err != nil {
		panic(fmt.Sprintf("Scanner contained an error: %s\n", err))
	}

	// how many?
	fmt.Printf("### %d measurements larger than the previous one.\n", counter)
}
