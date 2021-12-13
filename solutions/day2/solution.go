package main

import (
	"fmt"
	"io"
)

type command struct {
	Direction string
	Units     int
}

func readCommands() []command {
	var commands []command
	for {
		var cmd command
		_, err := fmt.Scanf("%s %d", &cmd.Direction, &cmd.Units)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		commands = append(commands, cmd)
	}
	return commands
}

func part1(cmds []command) int {
	var (
		depth              int
		horizontalPosition int
	)
	for _, cmd := range cmds {
		switch cmd.Direction {
		case "up":
			depth -= cmd.Units
		case "down":
			depth += cmd.Units
		case "forward":
			horizontalPosition += cmd.Units
		}
	}
	return depth * horizontalPosition
}

func part2(cmds []command) int {
	var (
		depth              int
		horizontalPosition int
		aim                int
	)
	for _, cmd := range cmds {
		switch cmd.Direction {
		case "up":
			aim -= cmd.Units
		case "down":
			aim += cmd.Units
		case "forward":
			horizontalPosition += cmd.Units
			depth += aim * cmd.Units
		}
	}
	return depth * horizontalPosition
}

func main() {
	cmds := readCommands()
	fmt.Println(part1(cmds))
	fmt.Println(part2(cmds))
}
