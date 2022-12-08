// https://adventofcode.com/2022/day/7
package main

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name        string
	files       []File
	directories []Directory
	parent      *Directory
}

func (d *Directory) addDirectory(name string) {
	d.directories = append(d.directories, Directory{name: name, parent: d})
}

func (d *Directory) addFile(name string, size int) {
	d.files = append(d.files, File{name: name, size: size})
}

func (d *Directory) size() int {
	sum := 0
	for _, file := range d.files {
		sum += file.size
	}
	for _, dir := range d.directories {
		sum += dir.size()
	}
	return sum
}

func (d *Directory) traverse(pad string) {
	fmt.Printf("%s- %s (dir)\n", pad, d.name)

	for _, dir := range d.directories {
		dir.traverse("  " + pad)
	}

	for _, file := range d.files {
		fmt.Printf("  %s- %s (file, size=%d)\n", pad, file.name, file.size)
	}
}

var sizeCap = 100_000

func sizesSum(d Directory) int {
	sum := 0
	size := d.size()
	if size < sizeCap {
		sum += size
	}

	for _, dir := range d.directories {
		sum += sizesSum(dir)
	}

	return sum
}

func main() {
	reader, err := os.Open("day07/input-challenge.txt")
	defer reader.Close()
	if err != nil {
		fmt.Printf("failed to open input file: %s\n", err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	rootDir := Directory{name: "/"}
	var cwd *Directory

	for scanner.Scan() {
		text := scanner.Text()
		if text == "$ ls" {
			continue
		}

		if text == "$ cd /" {
			cwd = &rootDir
			continue
		}

		if text == "$ cd .." {
			// keeping it simple, no validation
			cwd = cwd.parent
			continue
		}

		if text[:5] == "$ cd " {
			dir := text[5:]
			// keeping it simple, no validation
			for i, _ := range cwd.directories {
				if cwd.directories[i].name == dir {
					cwd = &cwd.directories[i]
					break
				}
			}
			continue
		}

		// ls

		// directory
		if text[:4] == "dir " {
			cwd.addDirectory(text[4:])

			continue
		}

		// file
		var name string
		var size int
		_, _ = fmt.Sscanf(text, "%d %s", &size, &name)

		cwd.addFile(name, size)
	}
	rootDir.traverse("")
	fmt.Println()
	fmt.Printf("Total sizes sum: %d\n", sizesSum(rootDir))
}
