package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const reportLength = 12

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	// get all data
	var diagnostics []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diagnostics = append(diagnostics, scanner.Text())
	}

	// calculate oxigen
	oxigen := calcRating(diagnostics, true)
	co2 := calcRating(diagnostics, false)

	fmt.Printf("Oxigen Generator Rating: %s\n", oxigen)
	fmt.Printf("CO2 scrubber rating Rating: %s\n", co2)
	fmt.Printf("Life Support Rating: %d\n", calcLifeSupportRating(oxigen, co2))
}

func calcLifeSupportRating(ox string, co2 string) int {
	g, err := strconv.ParseInt(ox, 2, 64)
	must(err)
	e, err := strconv.ParseInt(co2, 2, 64)
	must(err)

	return int(g * e)
}

func calcRating(diagnostics []string, isOxigenCalc bool) string {
	for i := 0; i < reportLength; i++ {
		if len(diagnostics) == 1 { // stop if only one number left
			break
		}
		b2k := calcBitToKeep(diagnostics, i, isOxigenCalc)
		diagnostics = filterDiagnostics(diagnostics, i, b2k)
	}
	return diagnostics[0]
}

// calcBitToKeep return 0 for bit 0 wins, or 1 for bit 1 wins
func calcBitToKeep(diagnostics []string, position int, isOxigenCalc bool) int {
	res := 0
	for _, v := range diagnostics {
		n, err := strconv.Atoi(fmt.Sprintf("%c", v[position]))
		must(err)
		switch n {
		case 0:
			res -= 1
		case 1:
			res += 1
		default:
			panic(fmt.Sprintf("no valid bit: %d", n))
		}
	}
	if isOxigenCalc && res < 0 {
		return 0
	} else if isOxigenCalc && res > 0 {
		return 1
	} else if isOxigenCalc && res == 0 {
		return 1
		// when CO2 calculations we take the fewer
	} else if !isOxigenCalc && res < 0 {
		return 1
	} else if !isOxigenCalc && res > 0 {
		return 0
	} else { //
		return 0
	}
}

// filterDiagnostics gets diagnostics with specific bit based in a bitPosition
func filterDiagnostics(diagnostics []string, bitPosition int, bitSearch int) []string {
	var res []string
	for _, v := range diagnostics {
		n, err := strconv.Atoi(fmt.Sprintf("%c", v[bitPosition]))
		must(err)
		if n == bitSearch {
			res = append(res, v)
		}
	}
	return res
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
