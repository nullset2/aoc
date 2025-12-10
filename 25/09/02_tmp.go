package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	points := make([]Point, 0)
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	maxx, maxy := 0, 0

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}

		parts := strings.Split(s, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		if x > maxx {
			maxx = x
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if y > maxy {
			maxy = y
		}
		points = append(points, Point{x: x, y: y})
	}

	grid := make([][]rune, maxy+1)
	for i := range grid {
		grid[i] = make([]rune, maxx+1)
	}

	for _, p := range points {
		grid[p.y][p.x] = '#'
	}

	A := len(points)

	for q, curr := range points {
		var next Point
		if q+1 < A {
			next = points[q+1]
		} else {
			next = points[0]
		}
		var left bool
		var above bool

		if next.x < curr.x {
			left = true
		}
		if next.y < curr.y {
			above = true
		}

		if left {
			for j := curr.x - 1; j > next.x; j-- {
				grid[curr.y][j] = 'X'
			}
		} else {
			for j := curr.x + 1; j < next.x; j++ {
				grid[curr.y][j] = 'X'
			}
		}

		if above {
			for j := curr.y - 1; j > next.y; j-- {
				grid[j][curr.x] = 'X'
			}
		} else {
			for j := curr.y + 1; j < next.y; j++ {
				grid[j][curr.x] = 'X'
			}
		}
	}

	for _, r := range grid {
		for _, q := range r {
			if q == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(q))
			}
		}

		fmt.Println()
	}

}
