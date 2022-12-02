// https://adventofcode.com/2022/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader, err := os.Open("day01/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	caloriesSum := 0
	top1 := 0
	top2 := 0
	top3 := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if caloriesSum >= top1 {
				top3, top2, top1 = top2, top1, caloriesSum
			} else if caloriesSum >= top2 {
				top3, top2 = top2, caloriesSum
			} else if caloriesSum > top3 {
				top3 = caloriesSum
			}
			caloriesSum = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("invalid input: %s\n", err)
			return
		}

		caloriesSum += calories
	}

	// deal with last elf
	if caloriesSum >= top1 {
		top3, top2, top1 = top2, top1, caloriesSum
	} else if caloriesSum >= top2 {
		top3, top2 = top2, caloriesSum
	} else if caloriesSum > top3 {
		top3 = caloriesSum
	}

	fmt.Println(top1 + top2 + top3)
}
