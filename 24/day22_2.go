package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func contains(s [][]int, t []int) bool {
	for _, slice := range s {
		if len(slice) != len(t) {
			continue
		}

		match := true
		for i := range slice {
			if slice[i] != t[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input22.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	dict := make(map[string]int, 0)

	for scanner.Scan() {
		var output []int
		var sequences [][]int
		s := scanner.Text()
		n, _ := strconv.Atoi(s)
		for i := 0; i < 2000; i++ {
			output = append(output, n)
			s1 := ((n * 64) ^ n) % 16777216
			s2 := ((int(math.Round(float64(s1 / 32)))) ^ s1) % 16777216
			s3 := (s2*2048 ^ s2) % 16777216
			n = s3
		}

		var digits []int
		for _, x := range output {
			digits = append(digits, x%10)
		}

		var visited [][]int

		for j := 0; j+4 < len(digits); j++ {
			var sequence []int

			sequence = append(sequence, digits[j+1]-digits[j])
			sequence = append(sequence, digits[j+2]-digits[j+1])
			sequence = append(sequence, digits[j+3]-digits[j+2])
			sequence = append(sequence, digits[j+4]-digits[j+3])
			sequences = append(sequences, sequence)

			str := strings.Trim(strings.Replace(fmt.Sprint(sequence), " ", ",", -1), "[]")
			if !contains(visited, sequence) { //only the first time it comes up for this monkey
				visited = append(visited, sequence)
				if _, ok := dict[str]; !ok {
					dict[str] = 0
				}
				dict[str] += digits[j+4]
			}

		}
	}

	max := -999999999
	for _, v := range dict {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)
}
