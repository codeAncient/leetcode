package main

import "fmt"

func main() {
	numArr := NewDifference([]int{1, 2, 3, 4, 5, 6})
	pNumArr := &numArr
	pNumArr.Increment(1, 3, 5)
	fmt.Println(pNumArr.Result())
}

type Difference struct {
	diff []int
}

func NewDifference(nums []int) Difference {

	diff := make([]int, len(nums)) //diff[0]==nums[0] diff[i]=nums[i]-nums[i-1] i>0
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return Difference{diff: diff}
}

func (this *Difference) Increment(i int, j int, val int) {
	this.diff[i] = this.diff[i] + val
	this.diff[j+1] = this.diff[j+1] - val
}

func (this *Difference) Result() []int {
	res := make([]int, len(this.diff))
	res[0] = this.diff[0]
	for i := 1; i < len(this.diff); i++ {
		res[i] = this.diff[i] + res[i-1]
	}
	return res
}
