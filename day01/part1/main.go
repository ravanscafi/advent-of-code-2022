// https://adventofcode.com/2022/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader, err := os.Open("day01/input-sample.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	caloriesSum := 0
	top := 0

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			if caloriesSum > top {
				top = caloriesSum
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
	if caloriesSum > top {
		top = caloriesSum
	}

	fmt.Println(top)
}
