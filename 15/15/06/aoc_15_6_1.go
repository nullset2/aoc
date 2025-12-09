package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}
	input, _ := os.Open("aoc_15_6.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		x1 := 0
		x2 := 0
		y1 := 0
		y2 := 0

		on := strings.Contains(s, "on")
		off := strings.Contains(s, "off")
		toggle := strings.Contains(s, "toggle")

		if on {
			x, _:= fmt.Sscanf(s, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Println(x)
		}

		if off {
			x, _:= fmt.Sscanf(s, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Println(x)
		}

		if toggle {
			x, _:= fmt.Sscanf(s, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			fmt.Println(x)
		}
		fmt.Println(x1, x2, y1, y2)

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if on {
					grid[i][j] = true
				}

				if off {
					grid[i][j] = false
				}

				if toggle {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}

	res := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				res++
			}
		}
	}

	fmt.Println(res)

}
