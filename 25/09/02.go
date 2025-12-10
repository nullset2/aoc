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

func eachCons(list []Point, n int) [][]Point {
	if n <= 0 || n > len(list) {
		return [][]Point{}
	}

	var result [][]Point
	for i := 0; i <= len(list)-n; i++ {
		window := list[i : i+n]
		result = append(result, append([]Point(nil), window...))
	}

	return result
}

func isValid(polygon [][]Point, a, b Point) bool {
	x1, y1 := a.x, a.y
	x2, y2 := b.x, b.y

	bx1, bx2 := min(x1, x2), max(x1, x2)
	by1, by2 := min(y1, y2), max(y1, y2)

	for _, side := range polygon {
		x3, y3 := side[0].x, side[0].y
		x4, y4 := side[1].x, side[1].y

		if max(x3, x4) > bx1 && min(x3, x4) < bx2 &&
			max(y3, y4) > by1 && min(y3, y4) < by2 {
			return false
		}
	}
	return true
}

func main() {
	points := make([]Point, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

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
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		points = append(points, Point{x: x, y: y})
	}

	n := len(points)
	maxArea := 0
	polygon := eachCons(points, 2)
	polygon = append(polygon, []Point{points[n-1], points[0]})
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := points[i]
			b := points[j]
			area := (1 + abs(a.x-b.x)) * (1 + abs(a.y-b.y))
			if area > maxArea && isValid(polygon, a, b) {
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)

}
