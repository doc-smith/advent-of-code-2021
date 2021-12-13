package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInitialState(scanner *bufio.Scanner) []uint8 {
	if scanner.Scan() {
		var state []uint8
		line := scanner.Text()
		for _, nStr := range strings.Split(line, ",") {
			n, err := strconv.ParseInt(nStr, 10, 8)
			if err != nil {
				panic(err)
			}
			state = append(state, uint8(n))
		}
		return state
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	panic("the input file is empty")
}

/* 3,4,3,1,2

     0: 0
     1: 1
     2: 1
     3: 2
     4: 1

   2,3,2,0,1

     0: 1
     1: 1
     2: 2
     3: 1
     4: 0

   1,2,1,6,0,8

     0: 0        0: 1
     1: 0        1: 2
     2: 0        2: 1
     3: 0        3: 0
     4: 0        4: 0
	 5: 0        5: 0
	 6: 1        6: 1
	 7: 0        7: 0
	 8: 1        8: 1
*/
func solve(initialState []uint8, numDays int) int64 {
	const numAgeStates = 9

	var currState [numAgeStates]int64
	for _, fishTimer := range initialState {
		currState[fishTimer]++
	}

	for day := 0; day < numDays; day++ {
		var nextState [numAgeStates]int64
		for i := 0; i < numAgeStates; i++ {
			nextState[i] = currState[(i+1)%numAgeStates]
		}
		nextState[6] += currState[0]
		currState = nextState
	}

	var fishCount int64
	for _, cnt := range currState {
		fishCount += cnt
	}
	return fishCount
}

func part1(initialState []uint8) int64 {
	return solve(initialState, 80)
}

func part2(initialState []uint8) int64 {
	return solve(initialState, 256)
}

func main() {
	initialState := readInitialState(
		bufio.NewScanner(os.Stdin),
	)
	fmt.Println(part1(initialState))
	fmt.Println(part2(initialState))
}
