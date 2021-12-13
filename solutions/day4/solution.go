package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readBalls(scanner *bufio.Scanner) []uint8 {
	if scanner.Scan() {
		var balls []uint8
		line := scanner.Text()
		for _, nStr := range strings.Split(line, ",") {
			n, err := strconv.ParseInt(nStr, 10, 8)
			if err != nil {
				panic(err)
			}
			balls = append(balls, uint8(n))
		}
		return balls
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	panic("the input file is empty")
}

type board = [][]uint8

func readBoard(scanner *bufio.Scanner) board {
	const boardSize = 5
	board := make([][]uint8, boardSize)
	for i := range board {
		board[i] = make([]uint8, boardSize)
		if scanner.Scan() {
			line := scanner.Text()
			for j, nStr := range strings.Fields(line) {
				n, err := strconv.ParseInt(nStr, 10, 8)
				if err != nil {
					panic(err)
				}
				board[i][j] = uint8(n)
			}
		} else {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
			panic("the board input is incomplete")
		}
	}
	return board
}

func mark(board board, ball uint8) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == ball {
				board[i][j] = 0
			}
		}
	}
}

func isWinner(board board) bool {
	for i := range board {
		boardSize := len(board[i])
		crossedInColumn := 0
		crossedInRow := 0
		for j := range board[i] {
			if board[i][j] == 0 {
				crossedInRow++
			}
			if board[j][i] == 0 {
				crossedInColumn++
			}
		}
		if crossedInRow == boardSize || crossedInColumn == boardSize {
			return true
		}
	}
	return false
}

func calcBoardScore(board board) int {
	score := 0
	for i := range board {
		for j := range board[i] {
			score += int(board[i][j])
		}
	}
	return score
}

func part1(balls []uint8, boards []board) int {
	for _, ball := range balls {
		for i := range boards {
			mark(boards[i], ball)
			if isWinner(boards[i]) {
				return calcBoardScore(boards[i]) * int(ball)
			}
		}
	}
	return -1
}

func part2(balls []uint8, boards []board) int {
	// there is no need to remember which boards are in play
	//   the order of boards is not important
	//   which means it's possible to remove the winning board in O(1)
	//     from the list of active boards (still in play):
	//       * exchange the winning boards with the last board
	//       * and decrement the number of boards that are still playable
	playableCount := len(boards)
	for _, ball := range balls {
		curBoard := 0
		for curBoard < playableCount {
			mark(boards[curBoard], ball)
			if isWinner(boards[curBoard]) {
				// this was the last board
				if playableCount == 1 {
					return calcBoardScore(boards[curBoard]) * int(ball)
				}
				playableCount--
				boards[curBoard], boards[playableCount] = boards[playableCount], boards[curBoard]
			} else {
				curBoard++
			}
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	balls := readBalls(scanner)
	var boards []board
	for {
		if scanner.Scan() {
			boards = append(boards, readBoard(scanner))
		} else {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
			break
		}
	}

	fmt.Println(part1(balls, boards))
	fmt.Println(part2(balls, boards))
}
