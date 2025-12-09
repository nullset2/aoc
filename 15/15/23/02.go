package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reg := map[rune]int{
		'a': 1,
		'b': 0,
	}
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	inst := make([]string, 0)
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		inst = append(inst, s)
	}

	i := 0

	for i >= 0 && i < len(inst) {
		in := inst[i]
		parts := strings.SplitN(in, " ", 2)
		p := parts[0]
		r := rune(parts[1][0])
		if p == "hlf" {
			reg[r] /= 2
		} else if p == "tpl" {
			reg[r] *= 3
		} else if p == "inc" {
			reg[r]++
		} else if p == "jmp" {
			offset, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			i += offset
			continue
		} else if p == "jie" && reg[r]%2 == 0 {
			off := strings.Split(parts[1], ", ")
			offset, err := strconv.Atoi(off[1])
			if err != nil {
				panic(err)
			}
			i += offset
			continue
		} else if p == "jio" && reg[r] == 1 {
			off := strings.Split(parts[1], ", ")
			offset, err := strconv.Atoi(off[1])
			if err != nil {
				panic(err)
			}
			i += offset
			continue
		}
		if reg[r] < 0 {
			reg[r] = 0
		}
		i++
	}

	fmt.Println(reg['a'])
	fmt.Println(reg['b'])

}
