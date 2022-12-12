// https://adventofcode.com/2022/day/10
package main

import (
	"bufio"
	"fmt"
	"os"
)

var regX = 1
var cycles = 0

func spriteVisible() bool {
	c := cycles % 40
	return c == regX || c == regX-1 || c == regX+1
}

func doCycle() {
	pixel := "."
	if spriteVisible() {
		pixel = "#"
	}
	fmt.Printf(pixel)
	cycles++
	if cycles%40 == 0 {
		fmt.Println()
	}
}

func noop() {
	doCycle()
}

func addx(value int) {
	doCycle()
	doCycle()
	regX += value
}

func main() {
	reader, err := os.Open("day10/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "noop" {
			noop()
			continue
		}

		var value int
		fmt.Sscanf(text, "addx %d", &value)

		addx(value)
	}
}
