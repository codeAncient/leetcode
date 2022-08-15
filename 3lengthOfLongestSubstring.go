package main

import (
	"fmt"
)

func main() {
	s := "中国的黄河中国的水开户行号哈哈哈"
	strArr := []rune(s)
	maxLen, left, right := lengthOfLongestSubstring(s)
	fmt.Println(maxLen)
	for _, char := range strArr[left:right] {
		fmt.Printf("%c", char)
	}
	fmt.Println("")
}

//最长不重复子串 右边窗口扩大，如果有当前元素重复则收缩左窗口直到
func lengthOfLongestSubstring(s string) (int, int, int) {
	window := make(map[rune]int)
	var left, right, res, resLeft, resRight int
	strArr := []rune(s) //字符串可能有中文，使用rune数组

	for right < len(strArr) {
		var c rune = strArr[right]
		right++ //扩大右边范围
		window[c]++
		for window[c] > 1 { //重复，缩小左边范围
			d := strArr[left]
			left++
			window[d]--
		}

		if right-left > res { //当前长度与历史长度比较，大则更改结果
			resLeft = left
			resRight = right
			res = right - left
		}
	}
	return res, resLeft, resRight
}
