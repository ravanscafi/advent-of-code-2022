// https://adventofcode.com/2022/day/12
package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	row, col int
}

func (p *point) adjacentPoints() []point {
	return []point{
		{p.row - 1, p.col},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
		{p.row, p.col + 1},
	}
}

type graphPoint struct {
	distance int
	point
}

type queue struct {
	elements []*graphPoint
}

func newQueue() *queue {
	return &queue{
		elements: make([]*graphPoint, 0),
	}
}

func (q *queue) push(elem graphPoint) {
	q.elements = append(q.elements, &elem)
}

func (q *queue) pop() graphPoint {
	elem := q.elements[0]
	q.elements[0] = nil
	q.elements = q.elements[1:]
	return *elem
}

func (q *queue) isEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	reader, err := os.Open("day12/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var start point
	var grid [][]rune
	row := 0

	for scanner.Scan() {
		text := scanner.Text()
		newRow := make([]rune, len(text))

		for col, r := range text {
			if r == 'S' {
				r = 'a'
			} else if r == 'E' {
				// start from 'E' and climb down searching for any 'a'
				start = point{row, col}
				r = 'z'
			}
			newRow[col] = r
		}

		grid = append(grid, newRow)
		row++
	}

	visitedPoints := []graphPoint{{0, start}}
	minDistance := 0
	q := newQueue()
	q.push(visitedPoints[0])

	for !q.isEmpty() {
		p := q.pop()
		for _, np := range p.adjacentPoints() {
			if np.row < 0 || np.col < 0 || np.row >= len(grid) || np.col >= len(grid[0]) {
				continue
			}
			isVisited := false
			for _, v := range visitedPoints {
				if np == v.point {
					isVisited = true
					break
				}
			}
			if isVisited {
				continue
			}
			if grid[p.row][p.col]-grid[np.row][np.col] > 1 {
				continue
			}
			if grid[np.row][np.col] == 'a' {
				minDistance = p.distance + 1
				fmt.Printf("Fewest steps required: %d\n", minDistance)
				return
			}
			nvp := graphPoint{p.distance + 1, np}
			visitedPoints = append(visitedPoints, nvp)
			q.push(nvp)
		}
	}
}
