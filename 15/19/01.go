package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func StaggeredReplace(input, from, to string) []string {
	if from == "" {
		return []string{input}
	}

	var results []string
	searchStart := 0

	for {
		idx := strings.Index(input[searchStart:], from)
		if idx == -1 {
			break
		}

		absIdx := searchStart + idx
		replaced := input[:absIdx] + to + input[absIdx+len(from):]
		results = append(results, replaced)
		searchStart = absIdx + len(from)
	}

	return results
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	transformations := make([][]string, 0)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)

		parts := strings.Split(s, " => ")
		transformations = append(transformations, parts)
	}

	line, _, _ := reader.ReadLine()
	molecule := string(line)
	m := make(map[string]int, 0)

	for _, t := range transformations {
		res := StaggeredReplace(molecule, t[0], t[1])

		for _, r := range res {
			m[r] = 1
		}
	}

	fmt.Println(len(slices.Collect(maps.Keys(m))))

}
