package main

import "fmt"

func main() {
	arr := []int{3, 5, 11, 11, 11, 56, 333}
	target := 333
	resIndex := right(arr, target)
	fmt.Println(resIndex)
}

//搜索最右边
func right(arr []int, target int) int {

	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target { //收缩左边，往右边找，因为要找最右边的
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid
		}
	}
	if arr[left-1] != target {
		return -1
	}
	return left - 1
}

//搜索最左边
func left(arr []int, target int) int {

	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target { //收缩右边，往左边找，因为要找最左边的
			right = mid
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid
		}
	}
	//数过大，left越界为len   数过小，不会发生越界l=0 r=0
	if left >= len(arr) {
		return -1
	}
	if arr[left] != target {
		return -1
	}
	return left
}

func solution(arr []int, left int, right int, target int) int {
	resIndex := -1
	mid := left + (right-left)/2

	if arr[mid] < target {
		left = mid + 1
	} else if arr[mid] > target {
		right = mid - 1
	} else if arr[mid] == target {
		resIndex = mid
		return resIndex
	}

	if left < 0 || right >= len(arr) {
		return resIndex
	}
	return solution(arr, left, right, target)
}
