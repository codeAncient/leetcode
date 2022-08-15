package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 10, 11, 19, 34, 41, 47, 55}
	target := 51
	fmt.Println(twoSum(arr, target))
}

func twoSum(numbers []int, target int) []int {
	res := []int{-1, -1}
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		fmt.Println(sum)
		if sum == target { //返回位置要以1开始
			res[0] = left + 1
			res[1] = right + 1
			return res
		} else if sum > target { //和大了，则右指针往左挪，和将变小
			right--
		} else if sum < target {
			left++
		}
	}
	return res
}
