package main

import (
	"fmt"
	"math"
)

func main() {
	str := "中国的美丽山河"
	t := "美山"
	fmt.Println(minWindow(str, t))
}

//最小覆盖子串 包含子串中所有字符的最小字符串
func minWindow(s string, t string) string {
	strArr := []rune(s)
	tArr := []rune(t)
	need := make(map[rune]int)
	window := make(map[rune]int)
	valid := 0
	for _, tStr := range tArr {
		need[tStr]++
	}
	var left, right, leftRes, rightRes int
	minLen := math.MaxInt
	for right < len(strArr) {

		str := strArr[right] //取得右边值

		if _, ok := need[str]; ok { //更新窗口
			window[str]++
			if window[str] == need[str] {
				valid++
			}
			//更新窗口后才可能改变valid 所以for循环放入if里面
			for valid == len(need) { //包含了所有子串中字符，可以收缩左边窗口了

				//收缩左窗口可以更新结果值
				if right-left < minLen {
					minLen = right - left
					leftRes = left
					rightRes = right
				}

				lstr := strArr[left]
				left++
				if _, ok := need[lstr]; ok {
					if window[lstr] == need[lstr] { //第一次减时不等了 valid--
						valid--
					}
					window[lstr]--
				}
			}

		}
		right++

	}
	if minLen == math.MaxInt {
		return ""
	}
	return string(strArr[leftRes : rightRes+1])
}

//0值放到数组末尾
func moveZeroes(nums []int) ([]int, int) {
	if len(nums) < 2 {
		return nums, len(nums)
	}
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 { //新探索值非0
			if slow != fast { //快慢指针不同，才有必要交换两个位置的值
				nums[slow] = nums[fast]
				nums[fast] = 0
			}
			slow++ //不等于0slow就往前走
		}
		fast++
	}
	return nums, slow - 1
}
