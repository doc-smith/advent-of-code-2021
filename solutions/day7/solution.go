package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/doc-smith/adventofcode2021/util/imath"
)

func readCrabPositions() []int {
	scanner := bufio.NewScanner(os.Stdin)
	var positions []int
	if scanner.Scan() {
		line := scanner.Text()
		for _, nStr := range strings.Split(line, ",") {
			x, err := strconv.Atoi(nStr)
			if err != nil {
				panic(err)
			}
			positions = append(positions, x)
		}
		return positions
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	panic("the input is empty")
}

// this is the naïve brutefoce solution (quadratic)
//   maybe rewrite it as DP?
func findBestCost(positions []int, distanceToCost func(int) int) int {
	var (
		minPosition = math.MaxInt
		maxPosition = math.MinInt
	)
	for _, p := range positions {
		if p > maxPosition {
			maxPosition = p
		}
		if p < minPosition {
			minPosition = p
		}
	}

	bestCost := math.MaxInt
	for cp := minPosition; cp <= maxPosition; cp++ {
		cost := 0
		for _, p := range positions {
			cost += distanceToCost(imath.Abs(p - cp))
		}
		if cost < bestCost {
			bestCost = cost
		}
	}
	return bestCost
}

func part1(positions []int) int {
	return findBestCost(positions, func(distance int) int {
		return distance
	})
}

func part2(positions []int) int {
	gauss := func(n int) int {
		return n * (n + 1) / 2
	}
	return findBestCost(positions, gauss)
}

func main() {
	positions := readCrabPositions()
	fmt.Println(part1(positions))
	fmt.Println(part2(positions))
}
