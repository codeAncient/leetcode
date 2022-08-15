package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5, 5, 5, 6, 7, 7, 8, 8, 8, 8, 9}
	fmt.Println(removeDuplicates(arr))
}

//原地去除数组中重复的元素 slow fast两个指针，fast去先探索，非重复元素则赋值给slow
func removeDuplicates(nums []int) (int, []int) {
	if len(nums) <= 1 {
		return len(nums), nums
	}
	slow := 0
	fast := 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] { //新探索值不重复
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	res := nums[:slow+1]
	return len(res), res
}
