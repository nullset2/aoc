package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		maxStr := findMaxSubsequence(b, 12)
		fmt.Println(maxStr)
		
		conv, _ := strconv.ParseInt(maxStr, 10, 64)
		res += conv
	}
	
	fmt.Println(res)
}

func findMaxSubsequence(s string, k int) string {
	n := len(s)
	result := strings.Builder{}
	startIdx := 0
	
	for i := 0; i < k; i++ {
		// How many digits we still need after this one
		remaining := k - i - 1
		
		// We can look up to position where we still have enough digits left
		maxPos := n - remaining
		
		// Find the maximum digit in the valid range
		maxDigit := byte('0')
		maxIdx := startIdx
		
		for j := startIdx; j < maxPos; j++ {
			if s[j] > maxDigit {
				maxDigit = s[j]
				maxIdx = j
			}
		}
		
		result.WriteByte(maxDigit)
		startIdx = maxIdx + 1
	}
	
	return result.String()
}
