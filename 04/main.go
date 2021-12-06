package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const gridSize = 5

func main() {
	file, err := os.Open("input")
	must(err)
	defer file.Close()

	bgNumbers, bgCards := loadBingoData(file)
	var winnerCard *bingoCard

Bingo:
	for _, n := range bgNumbers {
		for _, card := range bgCards {
			bingoCalled := card.CheckNumber(n)
			if bingoCalled {
				winnerCard = card
				break Bingo
			}
		}
	}

	fmt.Println("### SQUID BINGO ###")
	fmt.Printf("Bingo Numbers. %v\n", bgNumbers)
	fmt.Printf("Winner: %+v\n", winnerCard)
	fmt.Printf("\tScore: %+v\n", winnerCard.CalcScore())
}

func loadBingoData(file *os.File) ([]int, []*bingoCard) {
	var bgNumbers []int
	var bgCards []*bingoCard
	scanner := bufio.NewScanner(file)
	readLine := 0
	counter := 0
	for scanner.Scan() {
		rText := scanner.Text()
		readLine++
		if readLine == 1 {
			bgNumbers = readNumberList(rText, ",")
			continue
		}

		if len(removeNonNumbers(rText)) == 0 {
			continue
		}

		// new bingoCard
		if counter%gridSize == 0 {
			nbg := NewBingoCard()
			nbg.Numbers[counter%gridSize] = readNumberList(rText, " ")
			bgCards = append(bgCards, nbg)
			counter++
		} else { // add to last bingoCard
			bgCards[counter/gridSize].Numbers[counter%gridSize] = readNumberList(rText, " ")
			counter++
		}

	}
	return bgNumbers, bgCards
}

func removeNonNumbers(text string) string {
	reg, err := regexp.Compile("[^0-9]+")
	must(err)
	return reg.ReplaceAllString(text, "")
}

func readNumberList(text string, sep string) []int {
	// TODO: another option is to read 3 chars each time.
	var res []int
	tn := strings.Split(text, sep)
	for _, s := range tn {
		// this can be removed by doing it the way mentioned in upper TO-DO
		if len(removeNonNumbers(s)) < 1 {
			continue
		}
		n, err := strconv.Atoi(s)
		must(err)
		res = append(res, n)
	}
	return res
}

type bingoCard struct {
	Numbers       [][]int
	RowHit        []int
	ColumnHit     []int
	CalledNumbers []int
}

func (bc *bingoCard) CheckNumber(n int) (bingo bool) {
	bc.CalledNumbers = append(bc.CalledNumbers, n)
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if bc.Numbers[i][j] == n {
				bc.RowHit[i]++
				bc.ColumnHit[j]++
				if bc.RowHit[i] == 5 || bc.ColumnHit[j] == 5 {
					return true
				}
			}
		}
	}
	return false
}

// CalcScore sums all unmarked numbers then multiplies by the last called number
func (bc *bingoCard) CalcScore() int {
	return bc.sumUnmarkedNumbers() * bc.CalledNumbers[len(bc.CalledNumbers)-1]
}

func (bc *bingoCard) sumUnmarkedNumbers() int {
	sum := 0
	// sum all unmarked numbers
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if !contains(bc.CalledNumbers, bc.Numbers[i][j]) {
				sum += bc.Numbers[i][j]
			}
		}
	}
	return sum
}

func NewBingoCard() *bingoCard {
	bc := &bingoCard{}
	bc.Numbers = make([][]int, gridSize)
	for i := 0; i < 5; i++ {
		bc.Numbers[i] = make([]int, gridSize)
	}
	bc.RowHit = make([]int, gridSize)
	bc.ColumnHit = make([]int, gridSize)
	bc.CalledNumbers = []int{}
	return bc
}

func contains(l []int, n int) bool {
	for _, v := range l {
		if v == n {
			return true
		}
	}
	return false
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
