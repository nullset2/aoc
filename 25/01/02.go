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
		if len(s) == 0 {
			break
		}
		direction := s[0]
		size, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			for i := 0; i < size; i++{
				dial -= 1
				if dial < 0 {
					dial = 99
				}
				if dial == 0 {
					zeroes++
				}
			}
		} else if direction == 'R' {
			for i := 0; i < size; i++ {
				dial += 1
				if dial > 99 {
					dial = 0
				}
				if dial == 0 {
					zeroes++
				}
			}
		}
	}

	fmt.Println(zeroes)

}
