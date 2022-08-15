package main

import (
	"fmt"
)

func main() {
	arr := []byte("hollowa")
	fmt.Println(string(reverseString(arr)))
}
func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
	return s
}
