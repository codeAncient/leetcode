package main

import "fmt"

// import "fmt"

// func main() {
// 	arr := []int{5, 7, 9, 9, 2, 33, 21, 30}
// 	pq := &MinPQ{}
// 	for _, val := range arr {
// 		pq.Insert(val)
// 	}
// 	fmt.Println(pq.DelMin())
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lists := [][]int{{1, 47, 58}, {1, 3, 44}, {2, 6, 23, 76}}

	list := mergeKLists(lists)
	fmt.Println(traverse(list))

}
func traverse(list *ListNode) []int {
	if list == nil {
		return nil
	}
	res := make([]int, 0)
	p := list
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

//23合并k个升序链表
func mergeKLists(lists [][]int) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//将所有元素放入最小堆minPQ
	pq := &MinPQ{}
	pq.Insert(-1)
	for _, list := range lists {
		for _, val := range list {
			pq.Insert(val)
		}
	}
	fmt.Println(*pq)
	//准备链表 遍历最小堆 依次弹出最小元素
	dummy := &ListNode{Val: -1}
	p := dummy
	num := len(*pq)
	for num > 1 {
		min := pq.DelMin()
		p.Next = &ListNode{Val: min}
		p = p.Next
		num--
	}
	return dummy.Next
}

//优先级队列 父亲小于左右孩子，大则不配做父需要下沉 孩子比父节点小则不应该做孩子需要上浮.使用切片，0位置舍弃 所以位置>=1
type MinPQ []int

//x位置元素的父元素位置
func (m *MinPQ) Parent(x int) int {
	return x / 2
}

//x位置元素的左孩子位置
func (m *MinPQ) Left(x int) int {
	return x * 2
}

//x位置元素的右孩子位置
func (m *MinPQ) Right(x int) int {
	return x*2 + 1
}

//交互两个位置元素值
func (m *MinPQ) Swap(i int, j int) {
	q := *m
	q[i], q[j] = q[j], q[i]
}

//比较两个位置元素大小
func (m *MinPQ) Less(i int, j int) bool {
	q := *m
	return q[i] < q[j]
}

//上浮x位置元素 比较自己的父节点小则上浮
func (m *MinPQ) Swim(x int) {
	for x > 1 && m.Less(x, m.Parent(x)) { //x最多到堆顶
		m.Swap(m.Parent(x), x) //交换父子位置上的数值
		x = m.Parent(x)        //当前数的位置变为父亲的位置
	}
}

//下沉x位置元素 只要比左右孩子大就要下沉
func (m *MinPQ) Down(x int) {
	q := *m
	for m.Left(x) < len(q) { //x左子树如果超过边界就不可能再下沉了
		min := m.Left(x)                                    //假设左子树最大
		if m.Right(x) < len(q) && m.Less(m.Right(x), min) { //比较左右子树,得到最小子树
			min = m.Right(x)
		}
		//如果x小于子树 停止下沉
		if m.Less(x, min) {
			break
		}
		//否则下沉，交换
		m.Swap(min, x)
		x = min
	}
}

//删除顶部最小元素  交换顶部和最底部元素，并删除底部元素，顶部元素下沉到合适位置
func (m *MinPQ) DelMin() int {
	q := *m
	min := q[1]
	n := len(q)
	//fmt.Printf("min:%d \n", min)
	m.Swap(1, n-1)
	//fmt.Printf("未down未del:%v \n", q)
	*m = q[:n-1]
	m.Down(1)
	//fmt.Printf("down后:%v \n", q)
	return min
}

//插入元素 从尾部插入，进行上浮
func (m *MinPQ) Insert(v int) {
	*m = append(*m, v)
	m.Swim(len(*m) - 1)
}
