package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	set := make(map[int]int, 0)
	ranges := make([][]int, 0)
	ids := make([]int, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		parts := strings.Split(s, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, []int{start, end})
	}

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		id, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				set[id] = 1
			}
		}
	}

	fmt.Println(len(set))

}
