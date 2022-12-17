// https://adventofcode.com/2022/day/15
package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(sensor, beacon point) int {
	h := abs(sensor.x - beacon.x)
	v := abs(sensor.y - beacon.y)
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

	var sx, sy, bx, by int

	row := 2_000_000

	notBeacons := make(map[int]bool)
	beaconsAndSensors := make(map[int]bool)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := point{x: sx, y: sy}
		beacon := point{x: bx, y: by}
		if sensor.y == row {
			beaconsAndSensors[sensor.x] = true
		}
		if beacon.y == row {
			beaconsAndSensors[beacon.x] = true
		}
		dist := distance(sensor, beacon)
		size := dist

		if sensor.y <= row && sensor.y+dist >= row {
			size -= row - sensor.y
		} else if sensor.y > row && sensor.y-dist <= row {
			size -= sensor.y - row
		} else {
			continue
		}

		for x := sensor.x - size; x <= sensor.x+size; x++ {
			if !beaconsAndSensors[x] {
				notBeacons[x] = true
			}
		}
	}

	fmt.Printf("In the row y=%d, %d positions cannot contain a beacon", row, len(notBeacons))
}
