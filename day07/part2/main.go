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

func electDir(d Directory, reclaimSize int) Directory {
	electedDir := d
	electedDirSize := electedDir.size()

	for _, dir := range d.directories {
		tempDir := electDir(dir, reclaimSize)
		tempDirSize := tempDir.size()
		if tempDirSize >= reclaimSize && tempDirSize < electedDirSize {
			electedDir = tempDir
			electedDirSize = tempDirSize
		}
	}

	return electedDir
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

	size := rootDir.size()
	totalSpace := 70_000_000
	freeSpace := 30_000_000
	available := totalSpace - size
	toReclaim := freeSpace - available
	fmt.Printf("We need to reclaim %d\n", toReclaim)

	dir := electDir(rootDir, toReclaim)

	fmt.Printf("Best candidate for deletion is directory %q with size %d\n", dir.name, dir.size())
}
