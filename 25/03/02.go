package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	res := int64(0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	banks := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		banks = append(banks, s)
	}

	for _, b := range banks {
		n := len(b)
		max := math.MinInt64
		gen := combin.NewCombinationGenerator(n, 12)
		for gen.Next() {
			combination := gen.Combination(nil)
			sb := strings.Builder{}
			for _, x := range combination {
				sb.WriteByte(b[x])
			}
			str := sb.String()
			fmt.Println(str)
			conv, _ := strconv.Atoi(str)
			if conv > max {
				max = conv
			}
		}
		res += int64(max)
	}
	fmt.Println(res)
}
