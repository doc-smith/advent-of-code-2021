package main

import (
	"fmt"
	"io"
)

func readIntegers() []int {
	var integers []int

	var x int
	for {
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		} else {
			integers = append(integers, x)
		}
	}

	return integers
}

func part1(depthMap []int) int {
	increaseCount := 0
	for i := 1; i < len(depthMap); i++ {
		if depthMap[i] > depthMap[i-1] {
			increaseCount++
		}
	}
	return increaseCount
}

func part2(depthMap []int) int {
	const windowLen = 3

	windowSum := 0
	for i := 0; i < windowLen; i++ {
		windowSum += depthMap[i]
	}

	increaseCount := 0
	for i := windowLen; i < len(depthMap); i++ {
		nextWindowSum := windowSum - depthMap[i-windowLen] + depthMap[i]
		if nextWindowSum > windowSum {
			increaseCount++
		}
		windowSum = nextWindowSum
	}

	return increaseCount
}

func main() {
	depthMap := readIntegers()
	fmt.Println(part1(depthMap))
	fmt.Println(part2(depthMap))
}
