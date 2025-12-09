package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	reader := bufio.NewReader(input)
	res := 0
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		escaped := strconv.Quote(s)
		res += len(escaped) - len(s)
	}
	fmt.Println(res)
	return
}
