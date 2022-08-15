package main

import "fmt"

func main() {

	fmt.Println(corpFlightBookings([][]int{
		{1, 3, 1},
		{2, 4, 2},
		{5, 5, 1},
	}, 5))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	data := make([]int, n)
	numArr := NewDifference(data)
	pNumArr := &numArr
	for _, item := range bookings {
		pNumArr.Increment(item[0]-1, item[1]-1, item[2])
	}
	return pNumArr.Result()
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
	if i < len(this.diff) {
		this.diff[i] = this.diff[i] + val
	}
	if j < len(this.diff)-1 {
		this.diff[j+1] = this.diff[j+1] - val
	}

}

func (this *Difference) Result() []int {
	res := make([]int, len(this.diff))
	res[0] = this.diff[0]
	for i := 1; i < len(this.diff); i++ {
		res[i] = this.diff[i] + res[i-1]
	}
	return res
}
