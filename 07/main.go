package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabPositions := readCrabPositions("input")
	minPos, maxPos := getMinAndMaxCrabPosition(crabPositions)
	minFuelConsumption := 999999
	costEffectivePosition := 0

	for i := minPos; i <= maxPos; i++ {
		fc := calcFuelConsumption(crabPositions, i)
		if fc < minFuelConsumption {
			minFuelConsumption = fc
			costEffectivePosition = i
		}
	}
	fmt.Printf("Total fuel: %d for position %d\n", minFuelConsumption, costEffectivePosition)
	//printFuelConsumption(crabPositions, 2)
}

func calcFuelConsumption(crabPos []int, pos int) int {
	fuel := 0
	for _, p := range crabPos {
		fuel += int(math.Abs(float64(pos - p)))
	}
	return fuel
}

func getMinAndMaxCrabPosition(crabPos []int) (int, int) {
	min, max := crabPos[0], crabPos[0]
	for _, pos := range crabPos {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}
	return min, max
}

func readCrabPositions(path string) []int {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var pos []int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		ts := strings.Split(scanner.Text(), ",")
		for _, p := range ts {
			i, err := strconv.Atoi(p)
			must(err)
			pos = append(pos, i)
		}
	}
	return pos
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// printFuelConsumption helper function for debug
//func printFuelConsumption(crabPos []int, pos int) {
//	for i, po := range crabPos {
//		fuel := int(math.Abs(float64(pos - po)))
//		fmt.Printf("Crab %02d: Move from %02d to %02d: Fuel %d\n", i, po, pos, fuel)
//	}
//}
