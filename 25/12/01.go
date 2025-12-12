package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	res := 0
	var presents []int
	var areas []string
	var reqs [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		//Section 1
		if strings.Contains(s, "#") {
			count := 0
			for {
				s = scanner.Text()
				if len(s) == 0 {
					break
				}
				for _, c := range s {
					if c == '#' {
						count++
					}
				}
				scanner.Scan()
			}
			presents = append(presents, count)
		} else if strings.Contains(s, "x") {
			//section 2
			parts := strings.Split(s, ": ")
			var arr []int
			requirements := strings.Split(parts[1], " ")
			for _, r := range requirements {
				conv, err := strconv.Atoi(r)
				if err != nil {
					panic(err)
				}
				arr = append(arr, conv)
			}
			areas = append(areas, parts[0])
			reqs = append(reqs, arr)
		}
	}

	//check if available #s required can fit within the space
	for i := range areas {
		dimensions := areas[i]
		requirement := reqs[i]
		totalPresents := 0

		parts := strings.Split(dimensions, "x")
		m, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		available := m * n
		spaceNecessary := 0

		for j, r := range requirement {
			spaceNecessary += r * presents[j]
			totalPresents += r
		}

		if spaceNecessary <= available {
			res++
			continue
		}
	}
	fmt.Println(res)
}
