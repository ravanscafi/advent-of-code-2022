// https://adventofcode.com/2022/day/14
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type dimensions struct {
	min, max point
}

const startX = 500

func (d dimensions) relativeStartX() int {
	return startX - d.min.x
}

func newPointFromString(text string, d *dimensions) point {
	var x, y int
	fmt.Sscanf(text, "%d,%d", &x, &y)

	return newPoint(x, y, d)
}

func newPoint(x, y int, d *dimensions) point {
	if y > d.max.y {
		d.max.y = y
	}

	return point{
		x: x,
		y: y,
	}
}

// I could use [][]bool but then the print would look bad
type matrix [][]rune

func newMatrix(points []point, d dimensions) matrix {
	m := make(matrix, d.max.y-d.min.y+1)
	for y := range m {
		m[y] = make([]rune, d.max.x-d.min.x+1)
		for x := range m[y] {
			if y == 0 && x == d.relativeStartX() {
				m[y][x] = '+'
				continue
			}
			m[y][x] = '.'
		}
	}
	for _, p := range points {
		m[p.y][p.x-d.min.x] = '#'
	}
	return m
}

func (m matrix) print(d dimensions) {
	header1 := "    "
	header2 := "    "
	header3 := "    "
	body := ""
	for y := range m {
		body += fmt.Sprintf("%3d ", y)
		for x := range m[y] {
			if y == 0 {
				if x == 0 || x == d.max.x-d.min.x || x == d.relativeStartX() {
					xN := fmt.Sprintf("%d", x+d.min.x)
					header1 += string(xN[0])
					header2 += string(xN[1])
					header3 += string(xN[2])
				} else {
					header1 += " "
					header2 += " "
					header3 += " "
				}
			}

			body += fmt.Sprintf("%c", m[y][x])
		}
		body += fmt.Sprintln()
	}
	fmt.Println(header1)
	fmt.Println(header2)
	fmt.Println(header3)
	fmt.Println(body)
}

func (m matrix) addSand(d dimensions) bool {
	x := d.relativeStartX()

	if m[0][x] == 'o' {
		return false
	}

	for y := 0; y < d.max.y; y++ {
		if m[y+1][x] == '.' {
			continue
		}
		if m[y+1][x-1] == '.' {
			x--
			continue
		}
		if m[y+1][x+1] == '.' {
			x++
			continue
		}
		m[y][x] = 'o'
		return true
	}

	return false
}

func calculateLine(from, to point) []point {
	var points []point
	if from.x == to.x {
		start := from.y
		end := to.y

		if from.y > to.y {
			start = to.y
			end = from.y
		}

		for y := start; y <= end; y++ {
			points = append(points, point{x: from.x, y: y})
		}
		return points
	}

	start := from.x
	end := to.x

	if from.x < to.x {
		start = to.x
		end = from.x
	}

	for x := start; x >= end; x-- {
		points = append(points, point{x: x, y: from.y})
	}
	return points
}

func main() {
	reader, err := os.Open("day14/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var points []point

	d := dimensions{}

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), " -> ")

		for i := 0; i < len(path)-1; i++ {
			from := newPointFromString(path[i], &d)
			to := newPointFromString(path[i+1], &d)

			// wontfix: there are repeated points
			points = append(points, calculateLine(from, to)...)
		}
	}

	// add 2 more levels for the floor
	y := d.max.y + 2
	// adjust matrix x size to max base size
	d.min.x = 500 - y
	d.max.x = 500 + y
	// add floor points
	from := newPoint(d.min.x, y, &d)
	to := newPoint(d.max.x, y, &d)
	points = append(points, calculateLine(from, to)...)

	m := newMatrix(points, d)

	sandUnits := 0
	for m.addSand(d) {
		sandUnits++
	}

	//m.print(d)

	fmt.Printf("%d units of sand come to rest\n", sandUnits)
}
