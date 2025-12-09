package main

import "fmt"

func main() {
	s := "3113322113"
	for i := 0; i < 40; i++ {
		if len(s) == 0 {
			s = ""
		}
		curr := s[0]
		count := 1
		res := ""
		for i := 1; i < len(s); i++ {
			if s[i] == curr {
				count++
			} else {
				res += fmt.Sprintf("%d%d", count, curr-'0')
				curr = s[i]
				count = 1
			}
		}
		res += fmt.Sprintf("%d%d", count, curr-'0')
		s = res
	}
	fmt.Println(len(s))
	return
}
