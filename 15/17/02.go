package main

import (
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	res := 0
	input := `33
14
18
20
45
35
16
35
1
13
18
13
50
44
48
6
24
41
30
42`

	min := 20
	strings := strings.Split(input, "\n")
	fmt.Println(strings)
	containers := make([]int, 0)
	filtered := make([][]int, 0)
	for _, s := range strings {
		converted, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		containers = append(containers, converted)
	}

	n := len(containers)
	combinations := make([][]int, 0)

	for k := 0; k < n; k++ {
		combinations = append(combinations, combin.Combinations(n, k)...)
	}

	for _, c := range combinations {
		sum := 0
		for _, i := range c {
			sum += containers[i]
		}
		if sum == 150 {
			if len(c) < min {
				min = len(c)
			}
		}
	}


	for _, c := range combinations{
		if len(c) == min {
			filtered = append(filtered, c)
		}
	}

	for _, c := range filtered {
		sum := 0
		for _, i := range c {
			sum += containers[i]
		}
		if sum == 150 {
			res ++
		}
	}


	fmt.Println(res)
}
