package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
)

func main() {
	input, _ := os.Open("aoc_15_3.txt")
	defer input.Close()
	reader := bufio.NewReader(input)
	x1, y1 := 0, 0
	x2, y2 := 0, 0
	grid := make(map[string]int, 0)
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)

		curr := fmt.Sprintf("%d,%d", x1, y1)
		grid[curr]++
		curr = fmt.Sprintf("%d,%d", x2, y2)
		grid[curr]++

		for i, c := range s {
			if i%2 == 0 {
				if c == '>' {
					x1++
				} else if c == '^' {
					y1--
				} else if c == 'v' {
					y1++
				} else if c == '<' {
					x1--
				}
				curr = fmt.Sprintf("%d,%d", x1, y1)
				grid[curr]++
				continue
			} else {
				if c == '>' {
					x2++
				} else if c == '^' {
					y2--
				} else if c == 'v' {
					y2++
				} else if c == '<' {
					x2--
				}
				curr = fmt.Sprintf("%d,%d", x2, y2)
				grid[curr]++
				continue
			}
		}
	}

	fmt.Println(len(slices.Collect(maps.Keys(grid))))
}
