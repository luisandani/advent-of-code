package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	position := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		move := command[0]
		amount, err := strconv.Atoi(command[1])
		must(err)

		switch move {
		case "forward":
			position += amount
			depth += amount * aim
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	fmt.Printf("Position: %d Depth: %d Multpl: %d\n", position, depth, position*depth)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
