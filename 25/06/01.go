package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	operands := make([][]int, 0)
	operations := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}

		parts := regexp.MustCompile("[\\s]+").Split(s, -1)
		ints := make([]int, 0)
		if s[0] == '*' {
			operations = parts
			break
		}
		for _, p := range parts {
			conv, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			ints = append(ints, conv)
		}
		operands = append(operands, ints)
	}

	n := len(operands)
	m := len(operands[0])
	var total int64

	for j := 0; j < m; j++ {
		operation := operations[j]
		var curr int64
		if operation == "*" {
			curr = 1
		} else {
			curr = 0
		}
		for i := 0; i < n; i++ {
			if operation == "*" {
				curr *= int64(operands[i][j])
			} else {
				curr += int64(operands[i][j])
			}
		}
		total += curr
	}
	fmt.Println(total)

}
