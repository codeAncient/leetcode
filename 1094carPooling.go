package main

import "fmt"

func main() {

	fmt.Println(carPooling([][]int{
		{1, 0, 3},
		{2, 1, 4},
		{2, 3, 4},
	}, 3))
}

//0 <= fromi < toi <= 1000 最多有1001个站点， capacity为容量
//trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi
//区间下人或者上人，就是区间增减数据，要想不超载就是每个站点车上人数不大于容量
func carPooling(trips [][]int, capacity int) bool {
	data := make([]int, 1001)
	numArr := NewDifference(data)
	pNumArr := &numArr
	for _, item := range trips { //item[1]该站点就下人了，所以应该在上一个站点加人  站点从0编号
		pNumArr.Increment(item[1], item[2]-1, item[0])
	}
	res := pNumArr.Result()
	for _, val := range res {
		if val > capacity {
			return false
		}
	}
	return true
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
	if i < len(this.diff) { //i位置+val
		this.diff[i] = this.diff[i] + val
	}
	if j < len(this.diff)-1 { //j的下个位置需要-val，才会只在i-j位置上+val
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
