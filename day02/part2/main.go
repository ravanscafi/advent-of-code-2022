// https://adventofcode.com/2022/day/2
package main

import (
	"bufio"
	"fmt"
	"os"
)

var win = 6
var draw = 3
var lose = 0
var rock = 1
var paper = 2
var scissors = 3

var lookup = map[string]int{
	"A X": lose + scissors,
	"A Y": draw + rock,
	"A Z": win + paper,
	"B X": lose + rock,
	"B Y": draw + paper,
	"B Z": win + scissors,
	"C X": lose + paper,
	"C Y": draw + scissors,
	"C Z": win + rock,
}

func main() {
	reader, err := os.Open("day02/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	score := 0

	for scanner.Scan() {
		score += lookup[scanner.Text()]
	}

	fmt.Println(score)
}
