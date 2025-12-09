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
	x, y := 0, 0
	grid := make(map[string]int, 0)
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		curr := fmt.Sprintf("%d,%d", x, y)
		grid[curr]++
		for _, c := range s {
			if c == '>' {
				x++
			} else if c == '^' {
				y--
			} else if c == 'v' {
				y++
			} else if c == '<' {
				x--
			}
			curr = fmt.Sprintf("%d,%d", x, y)
			if _, ok := grid[curr]; !ok {
				grid[curr] = 1
				continue
			}
			grid[curr]++
		}
	}

	fmt.Println(len(slices.Collect(maps.Keys(grid))))
}
