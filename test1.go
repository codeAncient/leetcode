package main

import (
	"fmt"
	"reflect"
)

// 给你一个非递减数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 非递减：nums[j] >= nums[i], j > i
// 要求时间复杂度为 O(n)

// nums = [1,2,3,4,5]
// sorted(nums) = [1,4,9,16,25]

// nums = [-4,-1,0,3,10]
// sorted(nums) = [0,1,9,16,100]

// nums = [-7,-3,2,3,11]
// sorted(nums) = [4,9,9,49,121]

func sorted(nums []int) []int {
	// TODO 实现方法
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	i := right

	for left <= right {

		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]
		if leftVal > rightVal {
			res[i] = leftVal
			left++
		} else {
			res[i] = rightVal
			right--
		}

		i--
	}
	return res
}

func main() {
	cases := []*struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for i, c := range cases {
		got := sorted(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			fmt.Printf("case %d want %v but got %v\n", i, c.want, got)
		}
	}
}
