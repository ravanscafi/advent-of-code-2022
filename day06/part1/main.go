// https://adventofcode.com/2022/day/6
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader, err := os.Open("day06/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	buffer := ""
	markerSize := 4
	marker := -1

	for scanner.Scan() {
		marker++
		if len(buffer) < markerSize {
			buffer += scanner.Text()
			continue
		}

		found := true
		for i, r := range buffer {
			if index := strings.Index(buffer[i+1:], string(r)); index >= 0 {
				buffer = buffer[i:] // repeated char will be cut below
				found = false
				break
			}
		}

		if found {
			break
		}

		buffer = buffer[1:] + scanner.Text()
	}

	fmt.Println(buffer)
	fmt.Println(marker)
}
