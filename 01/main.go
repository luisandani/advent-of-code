package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open the file
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read line by line
	scan := bufio.NewScanner(file)
	depth := 0
	counter := 0
	for scan.Scan() {
		newDepth, err := strconv.Atoi(scan.Text())
		if err != nil {
			fmt.Printf("Error converting text to int. %s\n", err)
		}

		if depth == 0 {
			fmt.Printf("%d (N/A - no previous measurement)\n", newDepth)
		} else {
			op := "(decreased)"
			if newDepth > depth {
				op = "(increased)"
				counter++
			}
			fmt.Printf("%d %s\n", newDepth, op)
		}

		depth = newDepth
	}

	// check for errors
	if err := scan.Err(); err != nil {
		fmt.Printf("Scanner contained an error: %s\n", err)
	}

	// how many?
	fmt.Printf("### %d measurements larger than the previous one.", counter)
}
