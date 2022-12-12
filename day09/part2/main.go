// https://adventofcode.com/2022/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader, err := os.Open("day09/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	type point struct{ row, col int }
	positions := make(map[point]bool)
	knots := 9
	rope := make([]point, knots+1)
	var direction rune
	var steps int

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%c %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			switch direction {
			case 'L':
				rope[0].col--
			case 'R':
				rope[0].col++
			case 'U':
				rope[0].row--
			case 'D':
				rope[0].row++
			}

			for k := 1; k <= knots; k++ {
				p := point{
					row: rope[k].row - rope[k-1].row,
					col: rope[k].col - rope[k-1].col,
				}

				switch p {
				case point{2, 0}:
					rope[k].row--
				case point{-2, 0}:
					rope[k].row++
				case point{0, 2}:
					rope[k].col--
				case point{0, -2}:
					rope[k].col++
				case point{2, 1}, point{2, 2}, point{1, 2}:
					rope[k].row--
					rope[k].col--
				case point{2, -1}, point{2, -2}, point{1, -2}:
					rope[k].row--
					rope[k].col++
				case point{-2, 1}, point{-2, 2}, point{-1, 2}:
					rope[k].row++
					rope[k].col--
				case point{-2, -1}, point{-2, -2}, point{-1, -2}:
					rope[k].row++
					rope[k].col++
				}
			}

			positions[rope[knots]] = true
		}
	}

	fmt.Printf("Total positions visited by tail: %d\n", len(positions))
}
