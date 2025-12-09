package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ContainsRepeatedChar(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func main() {
	input, _ := os.Open("aoc_15_5.txt")
	defer input.Close()
	reader := bufio.NewReader(input)
	res := 0
	r1, err := regexp.Compile("(.*[aeiou].*){3,}")
	if err != nil {
		fmt.Println(err)
		return
	}
	r3, err := regexp.Compile("(ab|cd|pq|xy)")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)

		if r1.MatchString(s) && ContainsRepeatedChar(s) && !r3.MatchString(s) {
			res++
		}
	}
	fmt.Println(res)
}
