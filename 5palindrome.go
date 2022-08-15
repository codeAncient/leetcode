package main

import (
	"fmt"
)

func main() {
	s := "mooo"
	strArr := longestPalindrome(s)
	for _, char := range strArr {
		fmt.Printf("%c", char)
	}
	fmt.Println("\n-----")
}

func longestPalindrome(s string) []rune {
	var res []rune      //回文结果
	strArr := []rune(s) //转成rune切片 兼顾中文
	for i := 0; i < len(strArr); i++ {
		s1 := palindrome(strArr, i, i)   //奇数中心字符
		s2 := palindrome(strArr, i, i+1) //偶数中心字符
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(strArr []rune, l int, r int) []rune {
	for l >= 0 && r < len(strArr) && strArr[l] == strArr[r] {
		//如果没有越界，且左右相等满足回文，则扩大范围直到不满足条件
		l--
		r++
	}
	return strArr[l+1 : r] //切片右边是半开区间
}
