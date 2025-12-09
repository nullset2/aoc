package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	circuit := make(map[string]uint16)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)

		parts := strings.Split(s, " -> ")
		operands, assign := parts[0], parts[1]

		if strings.Contains(operands, "AND") {
			ops := strings.Split(operands, " AND ")
			res := circuit[ops[0]] & circuit[ops[1]]
			circuit[assign] = res
		} else if strings.Contains(operands, "OR") {
			ops := strings.Split(operands, " OR ")
			res := circuit[ops[0]] | circuit[ops[1]]
			circuit[assign] = res
		} else if strings.Contains(operands, "LSHIFT") {
			ops := strings.Split(operands, " LSHIFT ")
			n, _ := strconv.Atoi(ops[1])
			res := circuit[ops[0]] << n
			circuit[assign] = res
		} else if strings.Contains(operands, "RSHIFT") {
			ops := strings.Split(operands, " RSHIFT ")
			n, _ := strconv.Atoi(ops[1])
			res := circuit[ops[0]] >> n
			circuit[assign] = res
		} else if strings.Contains(operands, "NOT") {
			ops := strings.Replace(operands, "NOT ", "", -1)
			res := ^circuit[ops]
			circuit[assign] = res
		} else {
			n, _ := strconv.ParseUint(operands, 10, 16)
			circuit[assign] = uint16(n)
		}
	}
	fmt.Println(circuit["a"])

}
