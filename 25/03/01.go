package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	res := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	banks := make([][]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		bank := make([]int, 0)
		for _, c := range s {
			bank = append(bank, int(c-'0'))
		}
		banks = append(banks, bank)
	}

	for _, b := range banks {
		max := math.MinInt
		for i := range b {
			for _, c := range b[i+1:] {
				curr := 10*b[i] + c
				if curr > max {
					max = curr
				}
			}
		}
		fmt.Println(max)
		res += max
	}
	fmt.Println(res)
}
