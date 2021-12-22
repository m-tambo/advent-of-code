package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// first arg is the list of bingo numbers
	bingoNumbers := importBingoNums()

	// second arg is the file containing bingo boards
	boards := importBoards()

	var boardSums []int
	var winner *int
	var lastBingoNum *int
	winnerCount := 0

	// keep track of each board's total
	for i := 0; i < len(boards); i++ {
		boardSums = append(boardSums, sumAllElements(boards[i]))
	}

	// pull a number and then check for a winner
	for _, x := range bingoNumbers {
		for i := range boards {
			boards[i], boardSums[i] = markBoard(boards[i], boardSums[i], x)

			if checkForWin(boards[i]) {
				winner = &i
				lastBingoNum = &x
				break
			}
		}

		if winner != nil {
			break
		}
	}
	fmt.Println("Total score of first winner: ", boardSums[*winner] * *lastBingoNum)

	for _, x := range bingoNumbers {
		for i := range boards {
			boards[i], boardSums[i] = markBoard(boards[i], boardSums[i], x)

			if checkForWin(boards[i]) {
				winnerCount++
				if winnerCount == len(boards) {
					winner = &i
					lastBingoNum = &x
					break
				}
				boards[i] = make([]int, 25)
			}
		}
		if winnerCount == len(boards) {break}
	}
	fmt.Println("Total score of last winner: ", boardSums[*winner] * *lastBingoNum)

	// tests
	// fmt.Println("42 should equal ", sumAllElements([]int{3,15,0,2,22}))
}

func sumAllElements(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

// if a row or column contains 5 100s, it's a winner
func checkForWin(b []int) bool {
	if len(b) == 0 {
		return false
	}
	if sumAllElements([]int{b[0], b[1], b[2], b[3], b[4]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[5], b[6], b[7], b[8], b[9]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[10], b[11], b[12], b[13], b[14]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[15], b[16], b[17], b[18], b[19]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[20], b[21], b[22], b[23], b[24]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[0], b[5], b[10], b[15], b[20]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[1], b[6], b[11], b[16], b[21]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[2], b[7], b[12], b[17], b[22]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[3], b[8], b[13], b[18], b[23]}) > 499 {
		return true
	}
	if sumAllElements([]int{b[4], b[9], b[14], b[19], b[24]}) > 499 {
		return true
	}

	return false
}

// when a bingo number hits, mark that slot 100 and subtract the number from the board's total
// return the new board and the board's new total
func markBoard(board []int, boardTotal int, num int) ([]int, int) {
	newBoard := make([]int, len(board))
	_ = copy(newBoard, board)
	var newTotal int

	for i, a := range board {
		if a == num {
			newBoard[i] = 100
			newTotal = boardTotal - num
			return newBoard, newTotal
		}
	}

	return board, boardTotal
}

func importBingoNums() []int {
	nf := os.Args[1]
	numberFile, err := os.Open(nf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer numberFile.Close()

	var nums []int

	for {
		var num int
		_, err := fmt.Fscanf(numberFile, "%d", &num)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, num)
	}

	return nums
}

func importBoards() [][]int {
	f := os.Args[2]
	boardFile, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer boardFile.Close()

	var boards [][]int
	var board []int

	for {
		var element int
		_, err := fmt.Fscanf(boardFile, "%d", &element)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			if err.Error() == "unexpected newline" {
				boards = append(boards, board)
				board = []int{}
				continue
			}
			fmt.Println(err)
			os.Exit(1)
		}

		board = append(board, element)
	}
	return boards
}