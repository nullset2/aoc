package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Transpose(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]string, cols)

	for i := range result {
		result[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

func main() {
	res := int64(0)
	operands := make([][]string, 0)
	operations := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		if s[0] == '*' {
			operations = regexp.MustCompile("[\\s]+").Split(Reverse(strings.TrimSpace(s)), -1)
			break
		}
		parts := strings.Split(Reverse(s), "")
		operands = append(operands, parts)
	}

	transposed := Transpose(operands)

	fmt.Println(operations)
	fmt.Println(transposed)
	i := 0
	for j := 0; j < len(operations); j++ {
		o := operations[j]
		var curr int64
		if o == "+" {
			curr = 0
		} else {
			curr = 1
		}
		for i < len(transposed) {
			r := transposed[i]
			sb := strings.Builder{}
			allSpaces := true
			for _, n := range r {
				if n != " " {
					allSpaces = false
				}
			}
			if allSpaces {
				i++
				break
			}
			for _, n := range r {
				if n != " " {
					sb.WriteString(n)
				}
			}
			num := sb.String()
			if len(num) == 0 {
				break
			}
			conv, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			fmt.Println(conv, o)
			if o == "+" {
				curr += int64(conv)
			} else {
				curr *= int64(conv)
			}
			i++
		}
		res += curr
	}

	fmt.Println(res)

}
