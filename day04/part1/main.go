// https://adventofcode.com/2022/day/4
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader, err := os.Open("day04/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var elf1Min, elf1Max, elf2Min, elf2Max int
	overlaps := 0

	for scanner.Scan() {
		_, _ = fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elf1Min, &elf1Max, &elf2Min, &elf2Max)

		if (elf1Min >= elf2Min && elf1Max <= elf2Max) || (elf2Min >= elf1Min && elf2Max <= elf1Max) {
			overlaps++
		}
	}
	fmt.Println(overlaps)
}
