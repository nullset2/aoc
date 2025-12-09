package main

import (
	"fmt"
	"regexp"
)

func increment(s string) string {
	if len(s) == 0 {
		return "a"
	}

	ss := []rune(s)
	i := len(s) - 1

	for i >= 0 {
		if s[i] == 'z' {
			ss[i] = 'a'
			i--
		} else {
			ss[i] = ss[i] + 1
			break
		}
	}

	if i < 0 {
		ss = append([]rune{'a'}, ss...)
	}

	return string(ss)
}

func hasTwoDifferentPairs(s string) bool {
	pairs := make(map[rune]bool)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs[rune(s[i])] = true
			if len(pairs) >= 2 {
				return true
			}
		}
	}

	return false
}

func main() {
	password := "vzbxkghb"

	increasing := `abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz`
	charCheck := `i|o|l`

	inc := regexp.MustCompile(increasing)
	char := regexp.MustCompile(charCheck)

	for {
		if inc.MatchString(password) && !char.MatchString(password) && hasTwoDifferentPairs(password) {
			break
		}
		password = increment(password)
	}

	fmt.Println(password)
}
