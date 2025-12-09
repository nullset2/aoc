package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	KEY := "bgvyzdsv"
	curr := 0
	for {
		curr++
		c := fmt.Sprintf("%d", curr)
		hash := md5.Sum([]byte(KEY + c))
		md5String := hex.EncodeToString(hash[:])
		found := false
		for i := 0; i < 5; i++ {
			found = true
			if md5String[i] != '0' {
				found = false
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Println(curr)
}
