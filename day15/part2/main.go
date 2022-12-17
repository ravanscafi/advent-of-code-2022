// https://adventofcode.com/2022/day/15
package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func distance(sx, sy, bx, by int) int {
	h := abs(sx - bx)
	v := abs(sy - by)
	return h + v
}

func main() {
	reader, err := os.Open("day15/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	// By the problem constraints, there's only one possible point,
	// which means that by looking at the intersection of diagonal lines
	// of sensors, we can drastically reduce the amout of points to check.
	// Since there's no sensor with dist 1 of the beacons, if we find two
	// lines in the same direction with a distance of 1, it means that the
	// point is in the middle of those two parallel lines.
	// This goes for both positive lines (down->right) and negative
	// lines (down->left).
	// We need to identify both lines and intersect then to find the point.
	// Note: This solution doesn't work for the sample :(
	var positiveLines, negativeLines []int
	var sx, sy, bx, by int

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		d := distance(sx, sy, bx, by)
		positiveLines = append(positiveLines, sx-sy-d, sx-sy+d)
		negativeLines = append(negativeLines, sx+sy-d, sx+sy+d)
	}

	var positive, negative int
	size := len(positiveLines)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if abs(positiveLines[i]-positiveLines[j]) == 2 {
				positive = min(positiveLines[i], positiveLines[j]) + 1
			}
			if abs(negativeLines[i]-negativeLines[j]) == 2 {
				negative = min(negativeLines[i], negativeLines[j]) + 1
			}
		}
	}

	x := (positive + negative) / 2
	y := (negative - positive) / 2
	tuningFrequency := x*4_000_000 + y

	fmt.Printf("The beacon tuning frequency is %d\n", tuningFrequency)
}
