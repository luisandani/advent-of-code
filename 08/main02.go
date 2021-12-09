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
	file, err := os.Open("example")
	must(err)
	defer file.Close()

	var inputs []*line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 10 segments + | (delimiter) + 4 digit output value
		segments, output := getSegmentsAndOutputFromLine(scanner.Text())
		inputs = append(inputs, NewLine(segments, output))
	}

	for _, l := range inputs {
		rightMap := findingRightMapping(append(l.segments, l.output...))
		outputSum := calcOutputNumber(l.output, rightMap)
		fmt.Printf("%v: %d\n", l.output, outputSum)
	}

}

func calcOutputNumber(segments []string, rightMap [10]string) int {
	textRes := ""
	for _, seg := range segments {
		for i, s := range rightMap {
			if s == seg {
				textRes += strconv.Itoa(i)
			}
		}
	}
	res, err := strconv.Atoi(textRes)
	must(err)
	return res
}

func findingRightMapping(segments []string) (res [10]string) {
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
	}

	var charsOccurrences map[rune]int
	for _, s := range segments {
		for _, v := range s {
			charsOccurrences[v]++
		}
	}

	// missing 0, 2, 3, 5, 6, 9
	mapping := map[string]string{}

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

	// find number 9: in all segments with length 6 pick the one left with length 2 after removing number 4
	res[9] = getNumberNine(segments, res[4])
	// find "G" = 9 - 4 - mapAtA
	mapping["g"] = removeCharsFromString(removeCharsFromString(res[9], res[4]), mapping["a"])

	// finding numbers 6 = the remaining segment with length 2 from all with length 6
	// 		after removing number 4 + char mapped in A and that is not number 9
	res[6] = getNumberSix(segments, res[4], mapping["a"], res[9])

	// now I can calculate the "C" by substracting 8 -6
	mapping["c"] = removeCharsFromString(res[8], res[6])

	// "F" will be the 1 - the "C"
	mapping["f"] = removeCharsFromString(res[1], mapping["c"])

	// finding number 5 = the remaining empty string from all the ones with length 5
	//		after removing from it the number 6
	res[5] = getNumberFive(segments, res[6])

	// get number 0 = the segment left from all the ones with length 6 that are not segSix and segNine
	res[0] = getNumberZero(segments, res[6], res[9])

	// find "D" = 8 - 0
	mapping["d"] = removeCharsFromString(res[8], res[0])

	// get Number 3 = all segments with length 5 the one left with length 1 that is not segFive
	// 			after minus segNine
	res[3] = getNumberThree(segments, res[5], res[9])

	// b = 4 -3
	mapping["b"] = removeCharsFromString(res[4], res[3])

	// get Number 2 = 8 - "B" - "F"
	res[2] = removeCharsFromString(removeCharsFromString(res[8], mapping["b"]), mapping["f"])

	return
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

func getNumberThree(segments []string, segFive string, segNine string) string {
	for _, segment := range segments {
		if len(segment) != 5 {
			continue
		}
		if len(removeCharsFromString(segment, segNine)) == 0 && segment != segFive {
			return segment
		}
	}
	panic("couldn't find number Zero")
}

func getNumberZero(segments []string, segSix string, segNine string) string {
	for _, segment := range segments {
		if len(segment) != 6 {
			continue
		}
		if segment == segSix || segment == segNine {
			continue
		}
		return segment
	}
	panic("couldn't find number Zero")
}

func getNumberNine(segments []string, segFour string) string {
	for _, segment := range segments {
		if len(segment) != 6 {
			continue
		}
		step1 := removeCharsFromString(segFour, segment)
		if len(step1) == 2 {
			// got it!
			return segment
		}
	}
	panic("couldn't find number Nine")
}

func getNumberFive(segments []string, segSix string) string {
	for _, segment := range segments {
		if len(segment) != 5 {
			continue
		}
		step1 := removeCharsFromString(segment, segSix)
		if len(step1) == 0 {
			// got it!
			return segment
		}
	}
	panic("couldn't find number five")
}

func getNumberSix(segments []string, segFour string, mapAtA string, segNine string) string {
	for _, segment := range segments {
		if len(segment) != 6 {
			continue
		}
		step1 := removeCharsFromString(segment, segFour)
		step2 := removeCharsFromString(step1, mapAtA)
		if len(step2) == 2 && segment != segNine {
			// got it!
			return segment
		}
	}
	panic("couldn't find number six")
}

func removeCharsFromString(text string, remove string) string {
	for _, c := range remove {
		text = strings.ReplaceAll(text, string(c), "")
	}
	return text
}

func getOriginalMap() map[int]string {
	// a = 8 times, b=6, c=8, d=7, e=4, f=7, g=7
	return map[int]string{
		8: "abcdefg", // 7
		0: "abcefg",  // 6
		6: "abdefg",  // 6
		9: "abcdfg",  // 6
		2: "acdeg",   // 5
		3: "acdfg",   // 5
		5: "abdfg",   // 5
		4: "bcdf",    // 4
		7: "acf",     // 3
		1: "cf",      // 2
	}
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

func countSegments(signal string) int {
	segments := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0}
	count := 0
	for _, v := range signal {
		s := fmt.Sprintf("%c", v)
		segments[s]++
		if segments[s] == 1 { // first time we add, we increase the counter
			count++
		}
	}
	return count
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
