package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gonum/stat/combin"
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

	combinations := combin.Combinations(n, 2)
	max := math.MinInt

	for _, c := range combinations {
		p1 := points[c[0]]
		p2 := points[c[1]]

		area := (abs((p2.x - p1.x)) + 1) * (abs((p2.y - p1.y)) + 1)
		if area > max {
			max = area
		}
	}

	fmt.Println(max)

}
