package main

import "fmt"

func main() {
	numArr := Constructor([]int{1, 2, 3, 4, 5, 6})
	pNumArr := &numArr
	fmt.Println(pNumArr.SumRange(1, 3))
}

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {

	preSum := make([]int, len(nums)+1) //preSum[0]=0 preSum[i]代表前i个数的和
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	return NumArray{preSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	preSum := this.preSum
	return preSum[right+1] - preSum[left]
}
