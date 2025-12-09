package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mapIsContained(p, q map[string]int) bool {
	for k, v := range p {
		vv, ok := q[k]
		if !ok {
			return false
		}
		if k == "cats" || k == "trees" {
			if vv >= v {
				return false
			}
		} else if k == "pomeranians" || k == "goldfish" {
			if vv <= v {
				return false
			}
		} else if vv != v {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	aunts := make(map[int]map[string]int, 0)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		parts := strings.SplitN(s, ": ", 2)
		n, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Sue "))
		if err != nil {
			panic(err)
		}
		fields := strings.Split(parts[1], ", ")
		for _, f := range fields {
			parsed := strings.Split(f, ": ")
			x := parsed[0]
			y, err := strconv.Atoi(parsed[1])
			if err != nil {
				panic(err)
			}
			if _, ok := aunts[n]; !ok {
				aunts[n] = make(map[string]int, 0)
			}
			aunts[n][x] = y
		}
	}

	data := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for k, v := range aunts {
		if mapIsContained(v, data) {
			fmt.Println(k)
		}
	}
}
