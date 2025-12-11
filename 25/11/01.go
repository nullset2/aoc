package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res := 0
	machines := make(map[string][]string)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}

		parts := strings.Split(s, ": ")
		destinations := strings.Split(parts[1], " ")
		machines[parts[0]] = destinations
	}

	var stack []string
	curr := "you"
	stack = append(stack, curr)

	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == "out" {
			res++
		}

		if children, ok := machines[curr]; ok {
			stack = append(stack, children...)
		}
	}
	fmt.Println(res)
}
