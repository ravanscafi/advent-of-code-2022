// https://adventofcode.com/2022/day/11
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type id int
type item int

type monkey struct {
	id        id
	items     []item
	operation string
	test      int
	ifTrue    id
	ifFalse   id
	inspected int
}

func main() {
	reader, err := os.Open("day11/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	monkeys := make(map[id]*monkey)
	var current id = 0
	commomMultiple := 1

	for scanner.Scan() {
		if scanner.Text() == "" {
			current++
			continue
		}

		scanner.Scan()
		items := getItems(scanner.Text())

		scanner.Scan()
		operation := getOperation(scanner.Text())

		scanner.Scan()
		test := getTest(scanner.Text())
		commomMultiple *= test

		scanner.Scan()
		ifTrue := getIfTrue(scanner.Text())

		scanner.Scan()
		ifFalse := getIfFalse(scanner.Text())

		monkeys[current] = &monkey{
			id:        current,
			items:     items,
			operation: operation,
			test:      test,
			ifTrue:    ifTrue,
			ifFalse:   ifFalse,
		}
	}

	rounds := 10_000
	for r := 0; r < rounds; r++ {
		for m := id(0); int(m) < len(monkeys); m++ {
			for _, item := range monkeys[m].items {
				monkeys[m].inspected++
				item = doOperation(monkeys[m].operation, item, commomMultiple)

				dest := monkeys[m].ifFalse
				if int(item)%monkeys[m].test == 0 {
					dest = monkeys[m].ifTrue
				}
				monkeys[dest].items = append(monkeys[dest].items, item)
			}
			monkeys[m].items = []item{}
		}
	}

	top1 := 0
	top2 := 0
	for m := id(0); int(m) < len(monkeys); m++ {
		if monkeys[m].inspected >= top1 {
			top2 = top1
			top1 = monkeys[m].inspected
		} else if monkeys[m].inspected > top2 {
			top2 = monkeys[m].inspected
		}
		fmt.Printf("Monkey %d inspected items %d times.\n", int(m), monkeys[m].inspected)
	}
	monkeyBusiness := top1 * top2

	fmt.Printf("Monkey business level after %d rounds of stuff-slinging simian shenanigans: %d\n", rounds, monkeyBusiness)
}

func doOperation(operation string, i item, multiple int) item {
	fields := strings.Fields(strings.ReplaceAll(operation, "old", strconv.Itoa(int(i))))

	result := 0
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[2])
	switch fields[1] {
	case "+":
		result = a + b
	case "*":
		result = a * b
	}

	return item(result % multiple)
}

func getItems(text string) []item {
	itemsStr := strings.Split(text, ": ")[1]
	var items []item
	for _, s := range strings.Split(itemsStr, ", ") {
		i, _ := strconv.Atoi(s)
		items = append(items, item(i))
	}
	return items
}

func getOperation(text string) string {
	return strings.Split(text, "= ")[1]
}

func getTest(text string) int {
	var test int
	fmt.Sscanf(text, "  Test: divisible by %d", &test)
	return test
}

func getIfTrue(text string) id {
	var ifTrue id
	fmt.Sscanf(text, "    If true: throw to monkey %d", &ifTrue)
	return ifTrue
}

func getIfFalse(text string) id {
	var ifFalse id
	fmt.Sscanf(text, "    If false: throw to monkey %d", &ifFalse)
	return ifFalse
}
