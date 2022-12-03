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

	var rucksacks []string
	score := 0
	sum := 0

	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
		if len(rucksacks) < 3 {
			continue
		}

		var duplicatedItem rune
		for _, item := range rucksacks[0] {
			if strings.Contains(rucksacks[1], string(item)) && strings.Contains(rucksacks[2], string(item)) {
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
		rucksacks = []string{}
	}
	fmt.Println(sum)
}
