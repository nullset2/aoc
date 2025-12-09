package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	To   string
	Dist int
}

func main() {
	graph := make(map[string][]Entry)
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}

		s := string(line)

		parts := strings.Split(s, " = ")
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		points := strings.Split(parts[0], " to ")
		from, to := points[0], points[1]

		graph[from] = append(graph[from], Entry{To: to, Dist: dist})
	}

	min := math.MaxInt

	fmt.Println(min)

}
