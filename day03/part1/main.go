// https://adventofcode.com/2022/day/3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader, err := os.Open("day03/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	score := 0
	sum := 0

	for scanner.Scan() {
		rucksack := scanner.Text()
		size := len(rucksack) / 2
		compartment1, compartment2 := rucksack[0:size], rucksack[size:]

		var duplicatedItem rune
		for _, item := range compartment1 {
			if strings.Contains(compartment2, string(item)) {
				duplicatedItem = item
				break
			}
		}

		// between 'A' and 'Z'
		score = int(duplicatedItem) - 38

		// between 'a' and 'z'
		if duplicatedItem >= 96 {
			score -= 58
		}

		sum += score
	}
	fmt.Println(sum)
}
