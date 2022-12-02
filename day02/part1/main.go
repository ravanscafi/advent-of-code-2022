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
	"A X": draw + rock,
	"A Y": win + paper,
	"A Z": lose + scissors,
	"B X": lose + rock,
	"B Y": draw + paper,
	"B Z": win + scissors,
	"C X": win + rock,
	"C Y": lose + paper,
	"C Z": draw + scissors,
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
