// https://adventofcode.com/2022/day/13
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Comparison uint8

const (
	smaller   Comparison = iota
	equal     Comparison = iota
	bigger    Comparison = iota
	undefined Comparison = iota
)

type element interface {
	compare(el element) Comparison
}

type number struct {
	value int
}

type list struct {
	elements []element
}

func (n number) compare(el element) Comparison {
	switch typedEl := el.(type) {
	case number:
		return compareNumbers(n, typedEl)
	case list:
		newList := list{elements: []element{n}}
		return compareLists(newList, typedEl)
	}

	return undefined
}

func (l list) compare(el element) Comparison {
	switch typedEl := el.(type) {
	case number:
		newList := list{elements: []element{typedEl}}
		return compareLists(l, newList)
	case list:
		return compareLists(l, typedEl)
	}

	return undefined
}

func (l *list) addElement(e element) {
	l.elements = append(l.elements, e)
}

func main() {
	reader, err := os.Open("day13/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	indexSum := 0
	pairIndex := 1

	for scanner.Scan() {
		if scanner.Text() == "" {
			pairIndex++
			continue
		}

		left := createList(scanner.Text())

		scanner.Scan()
		right := createList(scanner.Text())

		if areListsOrdered(left, right) {
			indexSum += pairIndex
		}
	}

	fmt.Printf("The sum of the indices of the pairs is: %d\n", indexSum)
}

func areListsOrdered(left, right list) bool {
	return compareLists(left, right) == smaller
}

func compareLists(left, right list) Comparison {
	for i, l := range left.elements {
		if len(right.elements) <= i {
			return bigger
		}
		r := right.elements[i]

		if c := l.compare(r); c != equal {
			return c
		}
		continue
	}

	if len(left.elements) < len(right.elements) {
		return smaller
	}

	return equal
}

func compareNumbers(left, right number) Comparison {
	if left.value < right.value {
		return smaller
	}

	if left.value > right.value {
		return bigger
	}

	return equal
}

func createList(text string) list {
	l := list{}
	if text == "[]" {
		return l
	}
	openBrackets := 0
	buffer := ""

	for _, r := range text[1 : len(text)-1] {
		buffer += string(r)
		switch r {
		case '[':
			openBrackets++
		case ']':
			openBrackets--

			if openBrackets == 0 {
				l.addElement(createList(buffer))
				buffer = ""
			}
		case ',':
			if buffer == "," {
				buffer = ""
			}
			if buffer == "" || openBrackets > 0 {
				continue
			}

			l.addElement(createNumber(buffer[:len(buffer)-1]))
			buffer = ""
		}
	}

	if buffer != "" {
		l.addElement(createNumber(buffer))
	}

	return l
}

func createNumber(text string) number {
	n, _ := strconv.Atoi(text)
	return number{value: n}
}
