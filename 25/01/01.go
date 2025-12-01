package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	zeroes := 0
	dial := 50
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		if len(s) == 0 {
			break
		}
		direction := s[0]
		size, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			dial -= size
		} else if direction == 'R' {
			dial += size
		}
		dial = dial % 100
		if dial < 0 {
			dial = 100 - (dial * -1)
		}
		fmt.Println(dial)
		if dial == 0 {
			zeroes++
		}
	}

	fmt.Println(zeroes)

}
