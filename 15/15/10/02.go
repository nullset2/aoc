package main

import "fmt"

func main() {
    s := "3113322113"
    for i := 0; i < 50; i++ {
        s = lookAndSay(s)
    }
    fmt.Println(len(s))
}

func lookAndSay(s string) string {
    if len(s) == 0 {
        return ""
    }

    var result []byte
    curr := s[0]
    count := 1

    for i := 1; i < len(s); i++ {
        if s[i] == curr {
            count++
        } else {
            result = append(result, byte('0'+count), curr)
            curr = s[i]
            count = 1
        }
    }
    // Don't forget the last group
    result = append(result, byte('0'+count), curr)

    return string(result)
}
