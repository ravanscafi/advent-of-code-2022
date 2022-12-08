// https://adventofcode.com/2022/day/8
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader, err := os.Open("day08/input-sample.txt")
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

	visibleTrees := 0
	for r := range treeGrid {
		for c := range treeGrid[r] {
			t := treeGrid[r][c]

			if r == 0 || r == len(treeGrid)-1 ||
				c == 0 || c == len(treeGrid[r])-1 {
				visibleTrees++
				continue
			}

			// search left
			visible := true
			for i := r - 1; i >= 0; i-- {
				if treeGrid[i][c] >= t {
					visible = false
					break
				}
			}
			if visible {
				visibleTrees++
				continue
			}

			// search right
			visible = true
			for i := r + 1; i < len(treeGrid[r]); i++ {
				if treeGrid[i][c] >= t {
					visible = false
					break
				}
			}
			if visible {
				visibleTrees++
				continue
			}

			// search top
			visible = true
			for i := c - 1; i >= 0; i-- {
				if treeGrid[r][i] >= t {
					visible = false
					break
				}
			}
			if visible {
				visibleTrees++
				continue
			}

			// search down
			visible = true
			for i := c + 1; i < len(treeGrid); i++ {
				if treeGrid[r][i] >= t {
					visible = false
					break
				}
			}
			if visible {
				visibleTrees++
				continue
			}
		}
	}

	fmt.Printf("There are %d visible trees in the grid.\n", visibleTrees)
}
