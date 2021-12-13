package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/doc-smith/adventofcode2021/util/itertools"
	"github.com/doc-smith/adventofcode2021/util/queue"
)

func readHeightMap() [][]uint8 {
	var hm [][]uint8
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		heights := make([]uint8, len(line))
		for i := 0; i < len(line); i++ {
			heights[i] = line[i] - uint8('0')
		}
		hm = append(hm, heights)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return hm
}

func part1(hm [][]uint8) int {
	risk := 0
	for i := range hm {
		for j := range hm[i] {
			dis := [...]int{0, -1, 0, 1}
			djs := [...]int{-1, 0, 1, 0}
			isLow := itertools.ZipWhile(dis[:], djs[:], func(di int, dj int) bool {
				ni := i + di
				nj := j + dj
				if ni < len(hm) && ni >= 0 && nj < len(hm[i]) && nj >= 0 {
					if hm[ni][nj] <= hm[i][j] {
						return false
					}
				}
				return true
			})
			if isLow {
				risk += 1 + int(hm[i][j])
			}
		}
	}
	return risk
}

func gridBfs(
	visited [][]bool,
	si int,
	sj int,
	isReachable func(int, int) bool,
) int {
	type vertex struct {
		i int
		j int
	}

	q := queue.NewQueue()
	q.Enqueue(vertex{si, sj})

	componentSize := 0
	for !q.Empty() {
		v := q.Dequeue().(vertex)
		dis := [...]int{0, -1, 0, 1}
		djs := [...]int{-1, 0, 1, 0}
		itertools.Zip(dis[:], djs[:], func(di int, dj int) {
			ni := v.i + di
			nj := v.j + dj
			if isReachable(ni, nj) && !visited[ni][nj] {
				visited[ni][nj] = true
				q.Enqueue(vertex{ni, nj})
				componentSize++
			}
		})
	}
	return componentSize
}

func part2(hm [][]uint8) int {
	visited := make([][]bool, len(hm))
	for i := range visited {
		visited[i] = make([]bool, len(hm[i]))
	}

	var componentSizes []int

	for i := range hm {
		for j := range hm[i] {
			const wall = 9
			if hm[i][j] != wall && !visited[i][j] {
				sz := gridBfs(visited, i, j, func(vi int, vj int) bool {
					return vi >= 0 && vi < len(hm) &&
						vj >= 0 && vj < len(hm[i]) &&
						hm[vi][vj] != wall
				})
				componentSizes = append(componentSizes, sz)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(componentSizes)))
	res := 1
	for _, sz := range componentSizes[:3] {
		res *= sz
	}
	return res
}

func main() {
	heightMap := readHeightMap()
	fmt.Println(part1(heightMap))
	fmt.Println(part2(heightMap))
}
