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

	scan := bufio.NewScanner(file)
	counter := 0
	depth := 0
	for scan.Scan() {
		newDepth, err := strconv.Atoi(scan.Text())
		must(err)

		if depth > 0 && newDepth > depth {
			counter++
		}
		depth = newDepth
	}
	must(scan.Err())

	fmt.Printf("### %d measurements larger than the previous one.", counter)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
