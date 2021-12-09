package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	var inputs []*line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 10 segments + | (delimiter) + 4 digit output value
		segments, output := getSegmentsAndOutputFromLine(scanner.Text())
		inputs = append(inputs, NewLine(segments, output))
	}

	totalSum := 0
	for _, l := range inputs {
		rightMap := findingRightMapping(append(l.segments, l.output...))
		outputSum := calcOutputNumber(l.output, rightMap)
		totalSum += outputSum
		fmt.Printf("%v: %d\n", l.output, outputSum)
	}

	fmt.Printf("All output values: %d\n", totalSum)
}

func calcOutputNumber(segments []string, transMap map[string]string) int {
	correctMap := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}

	outputNumber := ""
	for _, segment := range segments {
		transSeg := translateSegment(segment, transMap)
		outputNumber += correctMap[transSeg]
	}
	res, err := strconv.Atoi(outputNumber)
	must(err)
	return res
}

func translateSegment(segment string, transMap map[string]string) string {
	res := ""
	for _, s := range segment {
		for k, v := range transMap {
			if v == string(s) {
				res += k
				break
			}
		}
	}
	return SortStringByCharacter(res)
}

func findingRightMapping(segments []string) map[string]string {
	res := make([]string, 10)
	for _, s := range segments {
		switch len(s) {
		case 2:
			res[1] = s
			continue
		case 3:
			res[7] = s
			continue
		case 4:
			res[4] = s
			continue
		case 7:
			res[8] = s
			continue
		}
	} // missing 0, 2, 3, 5, 6, 9

	charsOccurrences := make(map[rune]int)
	// i need to make a DISTINCT segments
	uniqueSegments := getUniqueSegments(segments)
	for _, s := range uniqueSegments {
		for _, r := range s {
			charsOccurrences[r]++
		}
	}

	mapping := make(map[string]string)
	// find "A" = 7 - 1
	mapping["a"] = removeCharsFromString(res[7], res[1])
	// first lets move with the ones we calculate by occurrences
	mapping["b"] = getCharByUniqueOccurrences(charsOccurrences, 6)      // only one with 6
	mapping["e"] = getCharByUniqueOccurrences(charsOccurrences, 4)      // only one with 4
	mapping["c"] = getCUsingOccurrences(charsOccurrences, mapping["a"]) // A & C only with 8
	// find "F" - 1 - C
	mapping["f"] = removeCharsFromString(res[1], mapping["c"])
	// find "G": 8 - 4 - mapA - mapE
	mapping["g"] = removeCharsFromString(removeCharsFromString(removeCharsFromString(res[8], res[4]), mapping["a"]), mapping["e"])
	// find "D" = just remove all the previous calculated
	mapping["d"] = calculateDUsingMappings(mapping)

	return mapping
}

func getUniqueSegments(segments []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range segments {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func calculateDUsingMappings(mapping map[string]string) string {
	res := "abcdefg"
	for _, v := range mapping {
		res = strings.ReplaceAll(res, v, "")
	}
	return res
}

func getCharByUniqueOccurrences(charsOccurrences map[rune]int, occurrences int) string {
	for r, count := range charsOccurrences {
		if count == occurrences {
			return string(r)
		}
	}
	panic("Couldn't find Char based on Occurrences")
}

func getCUsingOccurrences(charsOccurrences map[rune]int, s string) string {
	for r, count := range charsOccurrences {
		if count == 8 && s != string(r) {
			return string(r)
		}
	}
	panic("Couldn't find C based on Occurrences")
}

func removeCharsFromString(text string, remove string) string {
	for _, c := range remove {
		text = strings.ReplaceAll(text, string(c), "")
	}
	return text
}

type line struct {
	segments []string
	output   []string
}

func (l *line) String() string {
	return fmt.Sprintf("%v | %v\n", l.segments, l.output)
}

func NewLine(segments []string, output []string) *line {
	return &line{segments, output}
}

func getSegmentsAndOutputFromLine(textLine string) (patterns []string, output []string) {
	r := strings.Split(textLine, " | ")
	patterns = strings.Split(r[0], " ")
	for i := range patterns {
		patterns[i] = SortStringByCharacter(patterns[i])
	}
	output = strings.Split(r[1], " ")
	for i := range output {
		output[i] = SortStringByCharacter(output[i])
	}
	return
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
