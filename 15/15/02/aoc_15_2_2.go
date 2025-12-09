package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("aoc_15_2.txt")
	defer input.Close()

	reader := bufio.NewReader(input)
	res := 0
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		parts := strings.Split(s, "x")
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])
		p1 := 2 * (l + w)
		p2 := 2 * (w + h)
		p3 := 2 * (h + l)
		r := l * w * h

		perimeters := []int{p1, p2, p3}
		slices.Sort(perimeters)

		res += r + perimeters[0]
	}
	fmt.Println(res)
	return
}
