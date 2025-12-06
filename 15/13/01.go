package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	max := math.MinInt
	rules := make(map[string]map[string]int, 0)
	file, err := os.Open("mod_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		var orig, targ string
		var score int
		s = strings.TrimSuffix(s, ".")
		if strings.Contains(s, "gain") {
			fmt.Sscanf(s, "%s would gain %d happiness units by sitting next to %s.", &orig, &score, &targ)
		} else {
			fmt.Sscanf(s, "%s would lose %d happiness units by sitting next to %s.", &orig, &score, &targ)
			score = -1 * score
		}
		fmt.Println(orig, targ, score)

		if _, ok := rules[orig]; !ok {
			rules[orig] = make(map[string]int, 0)
		}
		rules[orig][targ] = score
	}
	n := len(rules)
	people := slices.Collect(maps.Keys(rules))
	permutations := combin.Permutations(n, n)

	for _, p := range permutations {
		score := 0
		for i, idx := range p {
			curr := people[idx]
			fmt.Print(curr + " ")
			//check left
			var l string
			if i == 0 {
				l = people[p[n-1]]
			} else {
				l = people[p[i-1]]
			}
			//check right
			var r string
			if i == n-1 {
				r = people[p[0]]
			} else {
				r = people[p[i+1]]
			}

			score += rules[curr][l] + rules[curr][r]
		}
		fmt.Println()
		fmt.Println(score)
		if score > max {
			max = score
		}
	}
	fmt.Println(max)
}
