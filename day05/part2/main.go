// https://adventofcode.com/2022/day/5
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader, err := os.Open("day05/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var lines []string

	// get first input part
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lines = append(lines, scanner.Text())
	}

	// identify the number of stacks
	stackNumbers := strings.Fields(lines[len(lines)-1])
	maxStack, _ := strconv.Atoi(stackNumbers[len(stackNumbers)-1])
	stacks := make([][]string, maxStack)

	// read crate lines and assign to `stacks` slices
	crates := lines[:len(lines)-1]
	for l := len(crates) - 1; l >= 0; l-- {
		stack := 0
		line := crates[l]
		for i := 0; i < len(line); i += 4 {
			if line[i+1] != ' ' {
				stacks[stack] = append(stacks[stack], string(line[i+1]))
			}
			stack++
		}
	}

	// perform the crane movements
	var quantity, from, to int
	for scanner.Scan() {
		_, _ = fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &quantity, &from, &to)
		from-- // slices are 0-indexed
		to--   // slices are 0-indexed
		lenF := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][lenF-quantity:lenF]...)
		stacks[from] = stacks[from][:lenF-quantity]
	}

	// identify the crate on top of each stack
	topCrates := ""
	for _, stack := range stacks {
		topCrates += stack[len(stack)-1]
	}
	fmt.Println(topCrates)
}
