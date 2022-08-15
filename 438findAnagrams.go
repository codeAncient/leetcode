package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s", string(20029))
	str := "中国的美丽山河中国的美丽山河中国的美丽山河"
	t := "美的"
	fmt.Println(checkInclusion(t, str))
}

//s1的字符串排列是否是s2的子串,s1字符串中的字符只能换顺序，但不允许有其他字符夹杂，要从s2中找这样的字符串
//满足字符情况，且必须是s1的长度就能保证这种情况
func checkInclusion(s1 string, s2 string) []int {
	res := make([]int, 0)
	strArr := []rune(s2)
	tArr := []rune(s1)
	need := make(map[rune]int)
	window := make(map[rune]int)
	valid := 0
	for _, tStr := range tArr {
		need[tStr]++
	}
	var left, right int
	for right < len(strArr) {

		str := strArr[right] //取得右边值

		if _, ok := need[str]; ok { //更新窗口
			window[str]++
			if window[str] == need[str] {
				valid++
			}
		}
		//只要满足子串长度就需要判断是否满足包含所有子串字符，包含则直接返回，不包含则收缩左边窗口
		if right-left+1 == len(tArr) {
			fmt.Printf("left:%v valid:%v need:%v window:%v\n", left, valid, len(need), window)
			if valid == len(need) {
				res = append(res, left)
			}
			//收缩左窗口
			lstr := strArr[left]
			left++
			if _, ok := need[lstr]; ok { //更新窗口
				if window[lstr] == need[lstr] {
					valid--
				}
				window[lstr]--
			}
		}
		right++

	}
	return res
}
