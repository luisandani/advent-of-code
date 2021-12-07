package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const daysToTrigger = 256

func main() {
	fishes := loadFishes("input")

	for day := 1; day <= daysToTrigger; day++ {
		// keep original number 0 will go to 6 and reproduce same amount at 8
		amountAt0 := fishes[0]
		for i := 0; i <= 7; i++ {
			fishes[i] = fishes[i+1]
		}
		fishes[8] = amountAt0
		fishes[6] += amountAt0
	}

	totalFish := 0
	for _, v := range fishes {
		totalFish += v
	}

	fmt.Printf("SUMMARY: After %02d day(s): %d\n", daysToTrigger, totalFish)
}

func loadFishes(path string) map[int]int {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var initFishes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initFishes = strings.Split(scanner.Text(), ",")
	}

	fishes := make(map[int]int, 9)
	for _, s := range initFishes {
		it, err := strconv.Atoi(s)
		must(err)
		if _, ok := fishes[it]; ok {
			fishes[it]++
		} else {
			fishes[it] = 1
		}
	}

	return fishes
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
