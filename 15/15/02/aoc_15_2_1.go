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
		a1 := l * w
		a2 := w * h
		a3 := h * l

		areas := []int{a1, a2, a3}
		slices.Sort(areas)

		res += 2*(a1+a2+a3) + areas[0]
	}
	fmt.Println(res)
	return
}
