package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(n int) bool {
	s := strings.Split(strconv.FormatInt(int64(n), 10), "")
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			val := i * j
			if val > max && isPalindrome(val) {
				max = val
			}
		}
	}
	fmt.Println(max)
}
