package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	var queue []int
	counter := 0
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		newDepth, err := strconv.Atoi(scan.Text())
		must(err)

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
	must(scan.Err())

	fmt.Printf("### %d measurements larger than the previous one.\n", counter)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
