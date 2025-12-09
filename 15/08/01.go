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
		escaped, err := strconv.Unquote(s)
		if err != nil {
			panic(err)
		}
		res += len(s) - len(escaped)
	}
	fmt.Println(res)
	return
}
