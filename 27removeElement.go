package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5, 5, 5, 6, 7, 7, 8, 8, 8, 8, 9}
	fmt.Println(removeElement(arr, 8))
}

//原地去除数组中=val的数, slow fast两个指针，fast去先探索，非val元素则赋值给slow
func removeElement(nums []int, val int) (int, []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val { //新探索值非val
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	res := nums[:slow]
	return len(res), res
}
