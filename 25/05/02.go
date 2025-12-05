package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := int64(0)
	ranges := make([][]int, 0)
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

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	result := make([][]int, 0)

	for _, i := range ranges {
		if len(result) == 0 || result[len(result)-1][1] < i[0] {
			result = append(result, i)
		} else {
			peek := result[len(result)-1]
			peek[1] = max(i[1], peek[1])
		}
	}

	for _, r := range result {
		res += int64(r[1] + 1 - r[0])
	}
	fmt.Println(res)
}
