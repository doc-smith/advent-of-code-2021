package main

import (
	"fmt"
	"io"

	"github.com/doc-smith/adventofcode2021/util/imath"
)

type point struct {
	Y, X int
}

type line struct {
	A, B point
}

func readLines() []line {
	var lines []line
	for {
		line := line{}
		_, err := fmt.Scanf(
			"%d,%d -> %d,%d",
			&line.A.X, &line.A.Y,
			&line.B.X, &line.B.Y,
		)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lines = append(lines, line)
	}
	return lines
}

func calcGridSize(lines []line) (int, int) {
	var (
		maxY int
		maxX int
	)

	for _, line := range lines {
		maxX = imath.Max(line.A.X, line.B.X, maxX)
		maxY = imath.Max(line.A.Y, line.B.Y, maxY)
	}

	return maxY + 1, maxX + 1
}

func solve(lines []line, ignoreDiagonal bool) int {
	n, m := calcGridSize(lines)

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}

	for _, line := range lines {
		currPoint := line.A
		deltaY := imath.Clamp(line.B.Y - line.A.Y)
		deltaX := imath.Clamp(line.B.X - line.A.X)
		if !ignoreDiagonal || deltaY == 0 || deltaX == 0 {
			for currPoint != line.B {
				grid[currPoint.Y][currPoint.X]++
				currPoint.Y += deltaY
				currPoint.X += deltaX
			}
			grid[currPoint.Y][currPoint.X]++
		}
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				res++
			}
		}
	}

	return res
}

func part1(lines []line) int {
	return solve(lines, true)
}

func part2(lines []line) int {
	return solve(lines, false)
}

func main() {
	lines := readLines()
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
