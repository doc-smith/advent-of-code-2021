package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	0: 6
	1: 2  *unique
	2: 5
	3: 5
	4: 4  *unique
	5: 5
	6: 6
	7: 3  *unique
	8: 7  *unique
	9: 6
*/

const (
	numDigits       = 10
	screenDigitSize = 4
)

type screenMeasurement struct {
	combinations [numDigits]string
	screen       [screenDigitSize]string
}

func readInput() []screenMeasurement {
	scanner := bufio.NewScanner(os.Stdin)
	var input []screenMeasurement
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != numDigits+screenDigitSize+1 {
			fmt.Println(fields)
			panic("the input is of wrong size")
		}
		measurement := screenMeasurement{}
		copy(measurement.combinations[:], fields[:numDigits])
		copy(measurement.screen[:], fields[len(fields)-screenDigitSize:])
		input = append(input, measurement)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return input
}

func part1(input []screenMeasurement) int {
	res := 0
	for _, measurement := range input {
		for _, count := range [...]int{2, 3, 4, 7} {
			for _, segments := range measurement.screen {
				if len(segments) == count {
					res++
				}
			}
		}
	}
	return res
}

func part2(input []screenMeasurement) int {
	return 0
}

func main() {
	input := readInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
