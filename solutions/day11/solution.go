package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/doc-smith/adventofcode2021/util/queue"
)

// there are always gridSize * gridSize octopi
const (
	gridSize  = 10
	maxEnergy = 9
)

func readGrid() [][]uint8 {
	var grid [][]uint8
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < gridSize; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			if len(line) != gridSize {
				panic("the grid line is corrupted")
			}
			grid = append(grid, make([]uint8, gridSize))
			for j := 0; j < gridSize; j++ {
				grid[i][j] = line[j] - uint8('0')
			}
		} else {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
			panic("the input is incomplete")
		}
	}
	return grid
}

func flash(grid [][]uint8, flashed [][]bool, si int, sj int) int {
	type vertex struct {
		i int
		j int
	}

	flashCount := 0

	q := queue.NewQueue()

	q.Enqueue(vertex{si, sj})
	flashed[si][sj] = true

	for !q.Empty() {
		v := q.Dequeue().(vertex)
		flashCount++
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				ni := v.i + di
				nj := v.j + dj
				inbound := ni >= 0 && ni < len(grid) &&
					nj >= 0 && nj < len(grid[ni])
				if inbound && !flashed[ni][nj] {
					grid[ni][nj]++
					if grid[ni][nj] > maxEnergy {
						flashed[ni][nj] = true
						q.Enqueue(vertex{ni, nj})
					}
				}
			}
		}
	}

	return flashCount

}

func makeStep(grid [][]uint8) int {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
		}
	}

	flashCount := 0

	flashed := make([][]bool, len(grid))
	for i := range flashed {
		flashed[i] = make([]bool, len(grid[i]))
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > maxEnergy && !flashed[i][j] {
				flashCount += flash(grid, flashed, i, j)
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if flashed[i][j] {
				grid[i][j] = 0
			}
		}
	}

	return flashCount
}

func part1(grid [][]uint8) int {
	const stepCount = 100
	flashCount := 0
	for i := 0; i < stepCount; i++ {
		flashCount += makeStep(grid)
	}
	return flashCount
}

func part2(grid [][]uint8) int {
	octopusCount := gridSize * gridSize
	step := 0
	for {
		step++
		flashCount := makeStep(grid)
		if flashCount == octopusCount {
			break
		}
	}
	return step
}

func copyGrid(grid [][]uint8) [][]uint8 {
	duplicate := make([][]uint8, len(grid))
	for i := range grid {
		duplicate[i] = make([]uint8, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func main() {
	grid := readGrid()
	fmt.Println(part1(copyGrid(grid)))
	fmt.Println(part2(copyGrid(grid)))
}
