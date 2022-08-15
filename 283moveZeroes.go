package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 0, 6, 42, 6, 0, 31, 5, 0, 0, 0, 12, 32}
	fmt.Println(moveZeroes(arr))
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
