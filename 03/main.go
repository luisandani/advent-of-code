package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const reportLength = 12 // 5 for the example

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	byts := make([]int, reportLength)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bn := strings.Split(scanner.Text(), "")
		for i, s := range bn {
			n, err := strconv.Atoi(s)
			must(err)
			switch n {
			case 0:
				byts[i] -= 1
			case 1:
				byts[i] += 1
			default:
				panic(fmt.Sprintf("no valid bit: %d", n))
			}
		}
	}

	// to binary
	gammaRate := calcGammaRate(byts)
	epsilonRate := calcEpsilonRate(byts)
	powerCons := calcPowerConsumption(gammaRate, epsilonRate)
	fmt.Printf("Power Consumption: %d. γ:%s | ε:%s\n", powerCons, gammaRate, epsilonRate)
}

func calcPowerConsumption(gamma string, epsi string) int {
	g, err := strconv.ParseInt(gamma, 2, 64)
	must(err)
	e, err := strconv.ParseInt(epsi, 2, 64)
	must(err)

	return int(g * e)
}

func calcGammaRate(b []int) string {
	g := ""
	for _, v := range b {
		switch {
		case v < 0:
			g += string('0')
		case v > 0:
			g += string('1')
		default:
			panic("that was not defined in the instructions to happen")

		}
	}
	return g
}

func calcEpsilonRate(b []int) string {
	e := ""
	for _, v := range b {
		switch {
		case v < 0:
			e += string('1')
		case v > 0:
			e += string('0')
		default:
			panic("that was not defined in the instructions to happen")

		}
	}
	return e
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
