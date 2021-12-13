package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return lines
}

func errorScore(line string) int {
	cost := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	bracketPairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	var stack []byte
	for i := 0; i < len(line); i++ {
		bracket := line[i]
		if closing, isOpen := bracketPairs[bracket]; isOpen {
			stack = append(stack, closing)
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == bracket {
				stack = stack[:len(stack)-1]
			} else {
				return cost[bracket]
			}
		}
	}

	return 0
}

func completionScore(line string) int {
	cost := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	bracketPairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	var stack []byte
	for i := 0; i < len(line); i++ {
		bracket := line[i]
		if closing, isOpen := bracketPairs[bracket]; isOpen {
			stack = append(stack, closing)
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == bracket {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}
	}

	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += cost[stack[i]]
	}
	return score
}

func part1(lines []string) int {
	score := 0
	for _, line := range lines {
		score += errorScore(line)
	}
	return score
}

func part2(lines []string) int {
	var scores []int
	for _, line := range lines {
		score := completionScore(line)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	lines := readLines()
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
