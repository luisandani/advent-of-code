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
	var winnersId []int
	var qCalledNumbers []int

LoserFound:
	for _, n := range bgNumbers {
		qCalledNumbers = append(qCalledNumbers, n)
		for _, card := range bgCards {
			bingoCalled := card.CheckNumber(n)
			if bingoCalled && !contains(winnersId, card.ID) {
				winnersId = append(winnersId, card.ID)
			}
			if (len(bgCards)) == len(winnersId) {
				card.Bingo = false
				break LoserFound
			}
		}
	}

	var loserCard *bingoCard
	for _, c := range bgCards {
		if !c.Bingo {
			loserCard = c
			break
		}
	}

	fmt.Println("### SQUID BINGO ###")
	fmt.Printf("Bingo Numbers. %v\n", bgNumbers)
	fmt.Printf("Fake Winner (Loser): %+v\n", loserCard)
	fmt.Printf("\tScore: %d\n", loserCard.CalcScore(qCalledNumbers))
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
			nbg.ID = counter / gridSize
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
	ID        int
	Numbers   [][]int
	RowHit    []int
	ColumnHit []int
	Bingo     bool
}

func (bc *bingoCard) CheckNumber(n int) (bingo bool) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if bc.Numbers[i][j] == n {
				bc.RowHit[i]++
				bc.ColumnHit[j]++
				if bc.RowHit[i] == 5 || bc.ColumnHit[j] == 5 {
					bc.Bingo = true
					return true
				}
			}
		}
	}
	return false
}

// CalcScore sums all unmarked numbers then multiplies by the last called number
func (bc *bingoCard) CalcScore(calledNumbers []int) int {
	return bc.sumUnmarkedNumbers(calledNumbers) * calledNumbers[len(calledNumbers)-1]
}

func (bc *bingoCard) sumUnmarkedNumbers(calledNumbers []int) int {
	sum := 0
	// sum all unmarked numbers
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if !contains(calledNumbers, bc.Numbers[i][j]) {
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
