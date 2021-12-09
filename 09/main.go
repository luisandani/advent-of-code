package main

import (
	"bufio"
	"os"
	"strconv"
)

const mapLength = 10 // 10 for example (100 for input)

func main() {
	file, err := os.Open("example")
	must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightMap := make([][]int, mapLength)
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			h, err := strconv.Atoi(v)
			must(err)
			heightMap
		}
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
