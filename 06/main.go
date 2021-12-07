package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const daysToTrigger = 80

func main() {
	fishes := loadFishes("input")

	for day := 1; day <= daysToTrigger; day++ {
		for _, fish := range fishes {
			reproduce := fish.TickDay()
			if reproduce {
				fishes = append(fishes, NewLanternFish(8))
			}
		}
	}
	fmt.Printf("SUMMARY: After %02d day(s): %d\n", daysToTrigger, len(fishes))
}

func loadFishes(path string) []*lanternfish {
	file, err := os.Open(path)
	must(err)
	defer file.Close()

	var initFishes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initFishes = strings.Split(scanner.Text(), ",")
	}

	var res []*lanternfish
	for _, s := range initFishes {
		it, err := strconv.Atoi(s)
		must(err)
		res = append(res, NewLanternFish(it))
	}

	return res
}

type lanternfish struct {
	IntTimer int
}

func (l *lanternfish) TickDay() (reproduce bool) {
	l.IntTimer--
	if l.IntTimer < 0 {
		l.IntTimer = 6
		reproduce = true
	}
	return reproduce
}

func (l *lanternfish) String() string {
	return fmt.Sprintf("%d", l.IntTimer)
}

func NewLanternFish(initTimer int) *lanternfish {
	return &lanternfish{initTimer}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
