// https://adventofcode.com/2022/day/8
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader, err := os.Open("day08/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	// parse text into a matrix
	var treeGrid [][]int
	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		treeGrid = append(treeGrid, make([]int, len(text)))
		for c, t := range text {
			treeGrid[row][c], _ = strconv.Atoi(string(t))
		}
		row++
	}

	bestScore := 0
	for r := range treeGrid {
		for c := range treeGrid[r] {
			score := 1
			t := treeGrid[r][c]

			if r == 0 || r == len(treeGrid)-1 ||
				c == 0 || c == len(treeGrid[r])-1 {
				continue
			}

			// search left
			sideScore := 0
			for i := r - 1; i >= 0; i-- {
				sideScore++
				if treeGrid[i][c] >= t {
					break
				}
			}
			score *= sideScore

			// search right
			sideScore = 0
			for i := r + 1; i < len(treeGrid[r]); i++ {
				sideScore++
				if treeGrid[i][c] >= t {
					break
				}
			}
			score *= sideScore

			// search top
			sideScore = 0
			for i := c - 1; i >= 0; i-- {
				sideScore++
				if treeGrid[r][i] >= t {
					break
				}
			}
			score *= sideScore

			// search down
			sideScore = 0
			for i := c + 1; i < len(treeGrid); i++ {
				sideScore++
				if treeGrid[r][i] >= t {
					break
				}
			}
			score *= sideScore

			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Printf("The highest scenic score possible is %d.\n", bestScore)
}
