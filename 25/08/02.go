package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/moorara/algo/unionfind"
)

type Point struct {
	x, y, z int
}

func Pow2(x int) int {
	return x * x
}

func Distance(p1, p2 Point) int {
	return int(float64(Pow2(p2.x-p1.x) + Pow2(p2.y-p1.y) + Pow2(p2.z-p1.z)))
}

func RemoveDupes(x [][]int) [][]int {
	set := make(map[int]int, 0)
	res := make([][]int, 0)
	for _, y := range x {
		if _, ok := set[y[0]]; !ok {
			res = append(res, y)
			set[y[0]] = 1
		}
	}
	return res
}

func main() {
	circuits := unionfind.NewQuickUnion(1000)
	points := make([]Point, 0)
	distances := make([][]int, 0)
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
		conv := []int{}
		for _, p := range parts {
			q, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}

			conv = append(conv, q)
		}

		points = append(points, Point{x: conv[0], y: conv[1], z: conv[2]})
	}

	for i, p1 := range points {
		for j, p2 := range points {
			if i != j {
				d := Distance(p1, p2)
				distances = append(distances, []int{d, i, j})
			}
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i][0] < distances[j][0]
	})
	distances = RemoveDupes(distances)

	for _, d := range distances {
		v1, _ := circuits.Find(d[1])
		v2, _ := circuits.Find(d[2])
		if v1 != v2 {
			circuits.Union(d[1], d[2])

			if circuits.Count() == 1 {
				fmt.Println(points[d[1]].x * points[d[2]].x)
			}
		}
	}

}
